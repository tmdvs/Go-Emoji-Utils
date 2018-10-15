package emoji

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
)

// Emoji - Struct representing Emoji
type Emoji struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	Descriptor string `json:"descriptor"`
}

// Emojis - Map of Emoji Runes as Hex keys to their description
var Emojis map[string]Emoji

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

	json.Unmarshal(byteValue, &Emojis)

	fmt.Print(len(Emojis))
}

// LookupEmoji - Lookup a single emoji definition
func LookupEmoji(emoji string) (Emoji, error) {
	// Convert our input string to UTF runes
	runes := []rune(emoji)

	// Build a slice of hex representations of each rune
	hexParts := []string{}
	for _, rune := range runes {
		hexParts = append(hexParts, fmt.Sprintf("%X", rune))
	}

	// Join the hex strings with a hypen - this is the key used in the emojis map
	hexKey := strings.Join(hexParts, "-")

	// If we have a definition for this string we'll return it,
	// else we'll return an error
	if emoji, ok := Emojis[hexKey]; ok {
		return emoji, nil
	}

	return Emoji{}, errors.New("No record for \"" + emoji + "\" could be found")
}

// LookupEmojis - Lookup definitions for each emoji in the input
func LookupEmojis(emoji []string) []interface{} {
	matches := []interface{}{}

	for _, emoji := range emoji {
		if match, err := LookupEmoji(emoji); err == nil {
			matches = append(matches, match)
		} else {
			matches = append(matches, err)
		}
	}

	return matches
}
