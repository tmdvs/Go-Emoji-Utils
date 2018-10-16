package emoji

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Convert a string to hex string representation of their Unicode Code Point value
func stringToHexKey(input string) (output string) {
	// Convert our input string to UTF runes
	runes := []rune(input)
	output = runesToHexKey(runes)
	return
}

// Convert a slice of runes to hex string representation of their Unicode Code Point value
func runesToHexKey(runes []rune) (output string) {
	// Build a slice of hex representations of each rune
	hexParts := []string{}
	for _, rune := range runes {
		hexParts = append(hexParts, fmt.Sprintf("%X", rune))
	}

	// Join the hex strings with a hypen - this is the key used in the emojis map
	output = strings.Join(hexParts, "-")
	return
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
