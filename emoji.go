package emoji

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
)

// Map of Emoji Runes as Hex keys to their description
var emojis map[string]string

// Unmarshal the emoji JSON into the Emojis map
func init() {
	// Work out where we are in relation to the caller
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	// Open the Emoji definition JSON and Unmarshal into map
	jsonFile, err := os.Open(path.Dir(filename) + "/data/emoji.json")
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	byteValue, e := ioutil.ReadAll(jsonFile)
	if e != nil {
		panic(e)
	}

	json.Unmarshal(byteValue, &emojis)
}

// DetectEmoji - Find all instances of emoji
func DetectEmoji(s string) map[string]int32 {

	// Convert our string to UTF runes
	runes := []rune(s)

	detectedEmojis := map[string]int32{}
	detectedModifiers := map[int]bool{}

	// Loop over each "word" in the string
	for index, rune := range runes {

		// If this index has been flaged as a modifier we do
		// not want to process it again
		if detectedModifiers[index] {
			continue
		}

		// Grab the initial hex value of this run
		hexKey := fmt.Sprintf("%X", rune)
		potentialMatches := emojis
		nextIndex := index + 1

		for {
			// Search the Emoji definitions map to see if we have
			// any matching results
			potentialMatches = searchEmojis(hexKey, potentialMatches)

			// We found a definitive match
			if len(potentialMatches) == 1 {
				break
			} else if len(potentialMatches) == 0 {
				// We didnt find anything, so we'll check if its a single rune emoji
				// Reset to original hexKey
				hexKey = fmt.Sprintf("%X", rune)
				if _, match := emojis[hexKey]; match {
					potentialMatches[hexKey] = emojis[hexKey]
				}

				// Definately no modifiers
				detectedModifiers = map[int]bool{}

				break
			} else {
				// Have we hit the last rune? If so we'll stop
				if nextIndex == len(runes) {
					// We need to return the match for this current key
					potentialMatches = map[string]string{
						hexKey: emojis[hexKey],
					}
					break
				}

				// We have more than one potential match so we'll add the
				// next UTF rune to the key and search again!
				hexKey = hexKey + "-" + fmt.Sprintf("%X", runes[nextIndex])
				detectedModifiers[nextIndex] = true
				nextIndex++
			}
		}

		for _, description := range potentialMatches {
			detectedEmojis[description]++
		}
	}

	// Return a map of Emojis and their counts
	return detectedEmojis
}

// Search our emoji definitions map for a key with a partial match
func searchEmojis(term string, list map[string]string) (results map[string]string) {

	results = map[string]string{}

	// Look for anything that has
	for key, value := range list {
		if strings.Index(key, term) == 0 {
			results[key] = value
		}
	}

	return results
}
