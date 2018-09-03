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

	// Loop over each "word" in the string
	for _, w := range strings.Split(s, " ") {
		r := []rune(w)

		// SPrint it as an array of hex parts
		s := fmt.Sprintf("%U", r)

		// Remove the UTF prefix and the slice boundaries
		s = strings.Replace(s, "U+", "", -1)
		s = strings.Replace(s, "[", "", -1)
		s = strings.Replace(s, "]", "", -1)
		s = strings.Replace(s, " ", "-", -1)

		// If it's an emoji count it up
		if _, ok := emojis[s]; ok {
			usedEmojis[emojis[s]]++
		}
	}

	// Return a map of Emojis and their counts
	return usedEmojis
}
