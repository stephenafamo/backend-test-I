package main

import (
	"bufio"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
)

func init() {
	setupConfig()
}

func main() {
	var wg sync.WaitGroup // so we can wait for our goroutine

	hashtags, minFollowers, maxFollowers := getFilterParams()

	wg.Add(1)
	go recordTweets(hashtags, minFollowers, maxFollowers)
	wg.Wait()
}

func recordTweets(hashtags string, minFollowers int, maxFollowers int) {

	sheet := getGoogleSheet()               // get the google sheet
	freeRow := getFreeRow(sheet)            // get which row to write to
	tweetStream := getTweetStream(hashtags) // get a stream of tweets

	for tweetInterface := range tweetStream {

		tweet, ok := tweetInterface.(*twitter.Tweet)
		if !ok {
			continue
		}

		followers := tweet.User.FollowersCount

		if followers >= minFollowers && followers <= maxFollowers {
			sheet.Update(freeRow, 0, tweet.User.ScreenName)
			sheet.Update(freeRow, 1, strconv.Itoa(tweet.User.FollowersCount))
			sheet.Update(freeRow, 2, tweet.Text)

			err := sheet.Synchronize()
			checkError(err)

			fmt.Printf("User: %#v Followers: %#v  TweetText: %#v Row: %#v \n", tweet.User.ScreenName, tweet.User.FollowersCount, tweet.Text, freeRow)
			freeRow++
		}
	}
}

func getFilterParams() (hashtags string, minFollowers int, maxFollowers int) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter search hashtags (comma separated): ")
	scanner.Scan()
	hashtags = scanner.Text()

	fmt.Printf("Enter minimum followers (default %#v): ", viper.GetInt("defaultMinFollowers"))
	scanner.Scan()
	minFollowers, _ = strconv.Atoi(scanner.Text())
	if minFollowers == 0 {
		minFollowers = viper.GetInt("defaultMinFollowers")
	}

	fmt.Printf("Enter maximum followers (default %#v): ", viper.GetInt("defaultMaxFollowers"))
	scanner.Scan()
	maxFollowers, _ = strconv.Atoi(scanner.Text())
	if maxFollowers == 0 {
		maxFollowers = viper.GetInt("defaultMaxFollowers")
	}

	if minFollowers > maxFollowers {
		panic("minimum followers must me less than maximum followers")
	}

	fmt.Printf("hashtags: %#v minFollowers: %#v  maxFollowers: %#v \n", hashtags, minFollowers, maxFollowers)

	return
}

func getTwitterClient() *twitter.Client {
	config := oauth1.NewConfig(viper.GetString("TwitterApiKey"), viper.GetString("TwitterApiSecret"))
	token := oauth1.NewToken(viper.GetString("TwitterAccessToken"), viper.GetString("TwitterAccessTokenSecret"))
	httpClient := config.Client(oauth1.NoContext, token)

	// twitter client
	client := twitter.NewClient(httpClient)

	return client
}

func getGoogleSheet() *spreadsheet.Sheet {
	data, err := ioutil.ReadFile(viper.GetString("GoogleApiCredentialsFile"))
	checkError(err)

	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	checkError(err)

	Gclient := conf.Client(context.TODO())

	service := spreadsheet.NewServiceWithClient(Gclient)

	spreadsheet, err := service.FetchSpreadsheet(viper.GetString("SpreadsheetId"))
	checkError(err)

	sheet, err := spreadsheet.SheetByIndex(uint(viper.GetInt("SheetIndex")))
	checkError(err)

	return sheet
}

func getTweetStream(hashtags string) chan interface{} {

	client := getTwitterClient()

	params := &twitter.StreamFilterParams{
		Track: []string{hashtags},
	}

	stream, err := client.Streams.Filter(params)
	checkError(err)

	return stream.Messages
}

func getFreeRow(sheet *spreadsheet.Sheet) (freeRow int) {

	freeRow = 0   // so we know which row is free
	rowIndex := 0 // keep track of searching our google sheet

	for _, row := range sheet.Rows {

		if freeRow != 0 {
			break // exit if the freerow has been updated
		}
		currentCell := 0

		for _, cell := range row {
			if currentCell > 2 {
				freeRow = rowIndex // if the cell first 3 cells are free, mark as a free row
			}
			if cell.Value != "" {
				continue // if a cell is not empty, move to the next row
			}
			currentCell++
		}
		rowIndex++
	}

	if freeRow == 0 {
		freeRow = rowIndex // In the case where all the rows in the google sheet are filled
	}

	return
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
