package auth

import (
	// other imports

	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// Credential to store the access data

type TwitterCredentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

type TwitterUserCredentials struct {
	OAuthToken       string
	OAuthTokenSecret string
}

func GetClient() (*twitter.Client, error) {
	// Pass in your consumer key (API Key) and your Consumer Secret (API Secret)
	creds, _ := generateCredentials()
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	// Pass in your Access Token and your Access Token Secret
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	user, _, err := client.Accounts.VerifyCredentials(verifyParams)

	if err != nil {
		return nil, err
	}
	log.Printf("user's account: \n %+v\n", user)
	return client, nil

}

func generateCredentials() (*TwitterCredentials, error) {
	tw_co_ke := os.Getenv("TWITTER_CONSUMER_KEY")
	tw_co_se := os.Getenv("TWITTER_CONSUMER_SECRET")
	tw_ac_to := os.Getenv("TWITTER_ACCESS_TOKEN")
	tw_ac_se := os.Getenv("TWITTER_ACCESS_SECRET")

	creds := TwitterCredentials{AccessToken: tw_ac_to, AccessTokenSecret: tw_ac_se, ConsumerKey: tw_co_ke, ConsumerSecret: tw_co_se}
	return &creds, nil
}
