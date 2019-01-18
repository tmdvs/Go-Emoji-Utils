package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tmdvs/Go-Emoji-Utils/utils"

	"github.com/tmdvs/Go-Emoji-Utils"

	"github.com/PuerkitoBio/goquery"
)

type lookup struct {
	Name string
	URL  string
}

func main() {
	fmt.Println("Updating Emoji Definition using Emojipedia…")

	// Grab the latest Apple Emoji Definitions
	res, err := http.Get("https://emojipedia.org/apple/")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	// Load the Apple Emoji HTML page into goquery so that we
	// can query the DOM
	doc, docErr := goquery.NewDocumentFromReader(res.Body)
	if docErr != nil {
		panic(docErr)
	}

	// Create a channel for lookups so that we can do this async
	lookups := make(chan lookup)

	go func() {
		// Find all emojis on the page
		doc.Find("ul.emoji-grid li").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the band and title
			emojiPage, _ := s.Find("a").Attr("href")
			title, _ := s.Find("img").Attr("title")

			fmt.Printf("Adding Emoji %d to lookups: %s - %s\n", i, title, emojiPage)

			// Add this specific emoji to the lookups to complete
			lookups <- lookup{
				Name: title,
				URL:  "https://emojipedia.org" + emojiPage,
			}
		})

		close(lookups)
	}()

	emojis := map[string]emoji.Emoji{}

	// Process a lookup
	for lookup := range lookups {
		fmt.Println("Looking up " + lookup.Name)

		// Grab the emojipedia page for this emoji
		res, err := http.Get(lookup.URL)
		if err != nil {
			fmt.Println(err)
			return
		}

		doc, docErr := goquery.NewDocumentFromReader(res.Body)
		if docErr != nil {
			panic(docErr)
		}

		emojiString, _ := doc.Find(".copy-paste input[type=text]").Attr("value")
		hexString := utils.StringToHexKey(emojiString)

		emojis[hexString] = emoji.Emoji{
			Key:        hexString,
			Value:      emojiString,
			Descriptor: lookup.Name,
		}

		fmt.Println(emojis[hexString])
	}

	s, _ := json.MarshalIndent(emojis, "", "\t")
	ioutil.WriteFile("data/emoji.json", []byte(s), 0644)
}
