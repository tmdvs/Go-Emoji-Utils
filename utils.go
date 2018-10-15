package emoji

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"
)

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
