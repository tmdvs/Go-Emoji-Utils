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
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

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
	usedEmojis := map[string]int32{}

	r := []rune(s)

	// Loop over each "word" in the string
	for i, c := range r {
		// Grab the initial hex value of this run
		hexKey := fmt.Sprintf("%X", c)

		// First lets check it it's been followed by a modifier
		nextIndex := i + 1

	modLoop:
		for {
			if nextIndex < len(r) {

				nextHexKey := hexKey + "-" + fmt.Sprintf("%X", r[nextIndex])

				for key := range emojis {
					if strings.Index(key, nextHexKey) >= 0 {
						nextIndex++
						hexKey = nextHexKey
						continue modLoop
					} else {
						break modLoop
					}
				}
			} else {
				break
			}
		}

		if _, ok := emojis[hexKey]; ok {
			usedEmojis[emojis[hexKey]]++
		}
	}

	// Return a map of Emojis and their counts
	return usedEmojis
}
