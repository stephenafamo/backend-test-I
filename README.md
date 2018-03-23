## Back-end Developer Test

### Devcenter Backend Developer Test I

The purpose of this test is not only to quickly gauge an applicant's abilities with writing codes, but also their approach to development.

Applicants may use whatever language they want to achieve the outcome.

## Task

Build a bot that extracts the following from people’s Twitter bio (on public/open accounts), into a Google spreadsheet:

* Twitter profile name 
* Number of followers

Target accounts using either of these criteria:

* Based on hashtags used
* Based on number of followers; Between 1,000 - 50,000

The bot is suppose to maintain a session and continously listen to the predefined hashtag

## Development Language

Go 1.10

## Application setup

1. Clone the repository
2. Create `config.json` from `config.json.example`
3. Setup your Twitter app as directed in the section below. Fill in the consumer key, consumer secret, access token and access token secret in your `config.json`
4. Setup your Google Service account, download the file called client_secret.json and place it at the root of the project and fill in the correct filename in the `config.json`.
5. Open the Google spreadsheet and share it with the service account id.
5. Install `dep`, the package manager for Go. See how [here](https://golang.github.io/dep/docs/installation.html).
6. Run `dep ensure` to install packages.
7. Run `go install` to build and install the application.
8. Run `backend-test-I`.

## Twitter App Setup

1. Create a Twitter account in case you don't have one. https://twitter.com/signup
2. Go to apps.twitter.com and click on 'Create New App ' button.
3. Fill out the details of the form correctly.
4. Then click on the ‘Create your Twitter application’ button.
5. Copy the consumer key, consumer secret, access token and access token secret to the `config.json`.

## Google App Setup
1. Create a service account for your Google project by visiting [this page](https://console.developers.google.com/projectselector/iam-admin/serviceaccounts).
2. Select a project or create a new project.
3. Click on "Create service account". Set any name you choose, set role as project->owner, and tick "furnish a new private key". When it is created, a file will be downloaded. Move it to the root of the project and set the correct filename in the `config.json`.
4. Open the chosen Google spreadsheet and share it with the Service Account ID (it looks like an email address).

## Demo

![screen shot](https://stephenafamo.com/blog/wp-content/uploads/2018/03/twitter-bot-demo.gif)
