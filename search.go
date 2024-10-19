package emoji

import (
	"errors"
	"strings"

	"github.com/tmdvs/Go-Emoji-Utils/utils"
)

// SearchResult - Occurrence of an emoji in a string
type SearchResult struct {
	Match       Emoji
	Occurrences int
	Locations   [][]int
}

// SearchResults - The result of a search
type SearchResults []SearchResult

// IndexOf - Check to see if search results contains a specific element
func (results SearchResults) IndexOf(result interface{}) int {
	for i, r := range results {
		if r.Match == result {
			return i
		}
	}

	return -1
}

// Find a specific emoji character within a string
func Find(emojiString string, input string) (result SearchResult, err error) {

	// Firstly we'll grab the emoji record for the emoji we're looking for
	var emoji Emoji
	if emoji, err = LookupEmoji(emojiString); err != nil {
		return // The emoji doesn't exist, return an error
	}

	// Find all of the emojis from in the input string
	allEmoji := FindAll(input)

	// Loop through emoji present in input and if any match the
	// emoji we're looking for we'll return the result
	for _, r := range allEmoji {
		if r.Match.Key == emoji.Key {
			result = r
			return
		}
	}

	// The emoji wasn't found, return an error
	return result, errors.New("Emoji was not found")
}

// FindAll - Find all instances of emoji
func FindAll(input string) (detectedEmojis SearchResults) {

	// Convert our string to UTF runes
	runes := []rune(input)

	// Any potential modifiers such as a skin tone/gender
	detectedModifiers := map[int]bool{}

	// Loop over each "word" in the string
	for index, r := range runes {

		// If this index has been flagged as a modifier we do
		// not want to process it again
		if detectedModifiers[index] {
			continue
		}

		// If the previous rune was a zero width joiner we'll skip this one
		// [Github issue](https://github.com/tmdvs/Go-Emoji-Utils/issues/12#issuecomment-1362747872)
		if index >= 1 {
			previousRune := []rune{runes[index-1]}
			if isRuneZeroWidthJoiner(previousRune) {
				continue
			}
		}

		// Grab the initial hex value of this run
		hexKey := utils.RunesToHexKey([]rune{r})

		// Ignore any basic runes, we'll get funny partials
		// that we don't care about
		if len(hexKey) < 2 {
			continue
		}

		previousKey := hexKey
		potentialMatches := Emojis
		nextIndex := index + 1

		for {
			// Search the Emoji definitions map to see if we have
			// any matching results
			potentialMatches = findEmoji(hexKey, potentialMatches)

			// We found a definitive match
			if len(potentialMatches) == 1 {
				break
			} else if len(potentialMatches) == 0 {
				// We didn't find anything, so we'll check if its a single rune emoji
				// Reset to original hexKey
				if _, match := Emojis[previousKey]; match {
					potentialMatches[previousKey] = Emojis[previousKey]
				}

				// Definitely no modifiers
				detectedModifiers = map[int]bool{}

				break
			} else {
				// Have we hit the last rune? If so we'll stop
				if nextIndex == len(runes) {
					// We need to return the match for this current key
					potentialMatches = map[string]Emoji{
						hexKey: Emojis[hexKey],
					}
					break
				}

				// We have more than one potential match so we'll add the
				// next UTF rune to the key and search again!
				previousKey = hexKey
				hexKey = hexKey + "-" + utils.RunesToHexKey([]rune{runes[nextIndex]})
				detectedModifiers[nextIndex] = true
				nextIndex++
			}
		}

		// Loop over potential matches and ensure we're not counting partials
		for key, e := range potentialMatches {
			if _, match := Emojis[key]; match {

				// How many runes does this emoji use
				emojiRuneLength := len(strings.Split(e.Key, "-"))

				// Have we already accounted for this match?
				if i := detectedEmojis.IndexOf(e); i != -1 {
					detectedEmojis[i].Occurrences++
					detectedEmojis[i].Locations = append(detectedEmojis[i].Locations, []int{index, index + emojiRuneLength})
				} else {
					detectedEmojis = append(detectedEmojis, SearchResult{
						Match:       e,
						Occurrences: 1,
						Locations: [][]int{
							{index, index + emojiRuneLength},
						},
					})
				}
			}
		}
	}

	// Return a map of Emojis and their counts
	return detectedEmojis
}

// Search an array of emoji definitions for a key with a partial match
func findEmoji(term string, list map[string]Emoji) (results map[string]Emoji) {
	results = map[string]Emoji{}

	// Look for anything that has
	for key, value := range list {
		if strings.Index(key, term) == 0 {
			results[key] = value
		}
	}
	return
}

func isRuneZeroWidthJoiner(r []rune) bool {
	return utils.RunesToHexKey(r) == "200D"
}
