package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/spf13/viper"
)

func init() {
	setupConfig()
}

func main() {
	config := oauth1.NewConfig(viper.GetString("TwitterApiKey"), viper.GetString("TwitterApiSecret"))
	token := oauth1.NewToken(viper.GetString("TwitterAccessToken"), viper.GetString("TwitterAccessTokenSecret"))
	httpClient := config.Client(oauth1.NoContext, token)

	// twitter client
	client := twitter.NewClient(httpClient)
	fmt.Printf("%v", client)
}
