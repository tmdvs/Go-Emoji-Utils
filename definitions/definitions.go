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

// Emoji - Struct representing Emoji
type Emoji struct {
	Key        string
	Value      string
	Descriptor string
}

// Emojis - Map of Emoji Runes as Hex keys to their description
var Emojis map[string]string

// Unmarshal the emoji JSON into the Emojis map
func init() {
	// Work out where we are in relation to the caller
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	// Open the Emoji definition JSON and Unmarshal into map
	jsonFile, err := os.Open(path.Dir(filename) + "/../data/emoji.json")
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	byteValue, e := ioutil.ReadAll(jsonFile)
	if e != nil {
		panic(e)
	}

	json.Unmarshal(byteValue, &Emojis)
}

// FindEmoji - Search our emoji definitions map for a key with a partial match
func FindEmoji(term string, list map[string]string) (results map[string]string) {

	results = map[string]string{}

	// Look for anything that has
	for key, value := range list {
		if strings.Index(key, term) == 0 {
			results[key] = value
		}
	}

	return results
}
