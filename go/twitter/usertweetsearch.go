package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	var (
		twUser        = flag.String("twuser", os.Getenv("TWITTER_USER"), "Twitter user name")
		twKey         = flag.String("twkey", os.Getenv("TWITTER_KEY"), "Twitter API Key")
		twSecret      = flag.String("twsecret", os.Getenv("TWITTER_SECRET"), "Twitter API Secret")
		twToken       = flag.String("twtoken", os.Getenv("TWITTER_TOKEN"), "Twitter API Token")
		twTokenSecret = flag.String("twtokensecret", os.Getenv("TWITTER_TOKEN_SECRET"), "Twitter API Token Secret")
	)
	flag.Parse()

	anaconda.SetConsumerKey(*twKey)
	anaconda.SetConsumerSecret(*twSecret)
	api := anaconda.NewTwitterApi(*twToken, *twTokenSecret)

	tweets, err := getTweets(api, *twUser)
	if len(tweets) != 0 {
		writeTweetsToFile(tweets, *twUser+"-fulltweets.json")
		writeTweetsToFile(filterTweets(tweets, "*"), *twUser+"-asterisktweets.json")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func getTweets(api *anaconda.TwitterApi, user string) ([]anaconda.Tweet, error) {
	v := url.Values{}
	v.Add("screen_name", user)
	v.Add("include_rts", "true")
	v.Add("exclude_replies", "false")
	v.Add("trim_user", "true")
	v.Add("count", "200")

	out := []anaconda.Tweet{}
	for {
		tweets, err := api.GetUserTimeline(v)
		if err != nil {
			// Still return all the tweets we had, just in case.
			return out, err
		}

		if len(tweets) == 0 {
			// We've gotten all the tweets the user has.
			break
		}
		lastTweet := tweets[len(tweets)-1]
		v.Set("max_id", strconv.FormatInt(lastTweet.Id-1, 10))
		out = append(out, tweets...)
	}

	return out, nil
}

func filterTweets(ts []anaconda.Tweet, term string) []anaconda.Tweet {
	out := []anaconda.Tweet{}
	for i := range ts {
		if strings.Contains(ts[i].Text, term) {
			out = append(out, ts[i])
		}
	}
	return out
}

func writeTweetsToFile(ts []anaconda.Tweet, file string) error {
	twJson, err := json.MarshalIndent(ts, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(file, twJson, 0644)
}
