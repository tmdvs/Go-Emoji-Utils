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

var emojis map[string]string

func init() {

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	jsonFile, err := os.Open(path.Dir(filename) + "/../data/emoji.json")
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

	for _, w := range strings.Split(s, " ") {
		r := []rune(w)
		s := fmt.Sprintf("%U", r)

		s = strings.Replace(s, "U+", "", -1)
		s = strings.Replace(s, "[", "", -1)
		s = strings.Replace(s, "]", "", -1)
		s = strings.Replace(s, " ", "-", -1)

		if _, ok := emojis[s]; ok {
			usedEmojis[emojis[s]]++
		}
	}

	return usedEmojis
}
