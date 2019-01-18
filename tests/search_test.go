package tests

import (
	"fmt"
	"testing"

	"github.com/tmdvs/Go-Emoji-Utils"
)

func BenchmarkSearch(b *testing.B) {

	fmt.Printf("Loaded %d emoji definitions\n", len(emoji.Emojis))

	b.ResetTimer()

	str := "👍🏻"
	for i := 0; i < b.N; i++ {
		emoji.FindAll(str)
	}
}

func BenchmarkComplexSearch(b *testing.B) {
	str := "This is a string 😄 🐷 with some 👍🏻🙈 emoji! 🐷 🏃🏿‍♂️"
	for i := 0; i < b.N; i++ {
		emoji.FindAll(str)
	}
}
