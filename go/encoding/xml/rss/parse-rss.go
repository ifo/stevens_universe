package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	feed, err := ioutil.ReadFile("feed.xml")
	if err != nil {
		log.Fatal(err)
	}

	var rss RSS
	err = xml.Unmarshal(feed, &rss)
	if err != nil {
		log.Fatal(err)
	}

	// print the links
	for _, i := range rss.Channel.Item {
		fmt.Println(i.Link)
	}
}

type RSS struct {
	Channel struct {
		Item []struct {
			Title       string `xml:"title"`
			Description string `xml:"description"`
			Link        string `xml:"link"`
			// Date won't parse, ignore for now.
			//PubDate     time.Time `xml:"pubDate"`
		} `xml:"item"`
	} `xml:"channel"`
}
