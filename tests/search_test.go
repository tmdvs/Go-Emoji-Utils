package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tmdvs/Go-Emoji-Utils"
)

func BenchmarkSearch(b *testing.B) {
	b.ResetTimer()
	str := "👍🏻"
	for i := 0; i < b.N; i++ {
		emoji.FindAll(str)
	}
}

func BenchmarkComplexSearch(b *testing.B) {
	b.ResetTimer()
	str := "This is a string 😄 🐷 with some 👍🏻 🙈 emoji! 🐷 🏃🏿‍♂️ 🥰"
	for i := 0; i < b.N; i++ {
		emoji.FindAll(str)
	}
}

func TestRemoveAllEmoji(t *testing.T) {

	str := "This is a string 😄 🐷 with some 👍🏻 🙈 emoji! 🐷 🏃🏿‍♂️ 🥰"

	matches := emoji.FindAll(str)
	totalUniqueEmoji := len(matches)

	assert.Equal(t, totalUniqueEmoji, 6, "There should be six different emoji")

	emojiRemoved := emoji.RemoveAll(str)
	assert.Equal(t, "This is a string with some emoji!", emojiRemoved, "There should be no emoji")

}
