// Package utils - Functions for converting between Hex, String, and Rune representations of emoji
package utils

import (
	"fmt"
	"strings"
)

// StringToHexKey - Convert a string to hex string representation of their Unicode Code Point value
func StringToHexKey(input string) (output string) {
	// Convert our input string to UTF runes
	runes := []rune(input)
	output = RunesToHexKey(runes)
	return
}

// RunesToHexKey - Convert a slice of runes to hex string representation of their Unicode Code Point value
func RunesToHexKey(runes []rune) (output string) {
	// Build a slice of hex representations of each rune
	hexParts := []string{}
	for _, rune := range runes {
		hexParts = append(hexParts, fmt.Sprintf("%X", rune))
	}

	// Join the hex strings with a hypen - this is the key used in the emojis map
	output = strings.Join(hexParts, "-")
	return
}
