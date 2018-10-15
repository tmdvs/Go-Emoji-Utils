package emoji

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strconv"
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

// Used to rebuild the raw emoji "values" in the emoji.json from their
// Key field which is the Hex representation of their Unicode Code Point value
func rebuildValuesFromHex() {
	for i, emoji := range Emojis {

		hexParts := strings.Split(emoji.Value, "-")

		s := []rune{}
		for _, p := range hexParts {
			n, _ := strconv.ParseUint(p, 16, 32)
			s = append(s, rune(n))
		}

		Emojis[i] = Emoji{
			Key:        emoji.Key,
			Value:      string(s),
			Descriptor: emoji.Descriptor,
		}
	}

	s, _ := json.MarshalIndent(Emojis, "", "\t")
	ioutil.WriteFile("data/emoji.json", []byte(s), 0644)

}
