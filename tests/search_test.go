package tests

import (
	"testing"

	emoji "github.com/tmdvs/Go-Emoji-Utils/search"
)

func BenchmarkSearch(b *testing.B) {
	str := "👍🏻"
	for i := 0; i < b.N; i++ {
		emoji.DetectEmoji(str)
	}
}

func BenchmarkComplexSearch(b *testing.B) {
	str := "This is a string 😄 🐷 with some 👍🏻🙈 emoji! 🐷 🏃🏿‍♂️"
	for i := 0; i < b.N; i++ {
		emoji.DetectEmoji(str)
	}
}
