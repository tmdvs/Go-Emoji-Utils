package tests

import (
	"testing"

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
	str := "This is a string 😄 🐷 with some 👍🏻🙈 emoji! 🐷 🏃🏿‍♂️"
	for i := 0; i < b.N; i++ {
		emoji.FindAll(str)
	}
}
