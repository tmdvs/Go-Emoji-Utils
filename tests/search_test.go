package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	emoji "github.com/tmdvs/Go-Emoji-Utils"
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

func TestContinuousEmoji(t *testing.T) {
	str := "abc🙏🙏🙏🙏🙏"

	matches := emoji.FindAll(str)
	totalUniqueEmoji := len(matches)

	assert.Equal(t, totalUniqueEmoji, 1, "There should be 1 unique emoji")

	emojiRemoved := emoji.RemoveAll(str)
	assert.Equal(t, "abc", emojiRemoved, "There should be no emoji")
}

func TestRemoveAllEmojiChinese(t *testing.T) {

	str := "起坎特在🇫🇷队的作用更      哈哈哈"

	matches := emoji.FindAll(str)
	totalUniqueEmoji := len(matches)

	assert.Equal(t, totalUniqueEmoji, 1, "There should be one emoji")

	emojiRemoved := emoji.RemoveAll(str)
	assert.Equal(t, "起坎特在队的作用更 哈哈哈", emojiRemoved, "There should be no emoji")

}

func TestRemoveAllEmojiChineseEnglishMixed(t *testing.T) {

	str := "wo🤮🤧武斌💁ello a武斌 g😇 🤠ood peo👌🎍😍ello"

	matches := emoji.FindAll(str)
	totalUniqueEmoji := len(matches)

	assert.Equal(t, totalUniqueEmoji, 8, "There should be one emoji")

	emojiRemoved := emoji.RemoveAll(str)
	assert.Equal(t, "wo武斌ello a武斌 g ood peoello", emojiRemoved, "There should be no emoji")

}

func TestRemoveAllEmojiJapanese(t *testing.T) {

	str := "被害者は深刻な影響を🤮🤧受けるにもか💁かわらず、被害だと😇 🤠認識できるま👌🎍😍で時間がかかり"

	matches := emoji.FindAll(str)
	totalUniqueEmoji := len(matches)

	assert.Equal(t, totalUniqueEmoji, 8, "There should be one emoji")

	emojiRemoved := emoji.RemoveAll(str)
	assert.Equal(t, "被害者は深刻な影響を受けるにもかかわらず、被害だと 認識できるまで時間がかかり", emojiRemoved, "There should be no emoji")

}

func TestRemoveAllEmojiKorean(t *testing.T) {

	str := "포기하고 싶은 순🤮간들 바💁로 그 순간   🤠빨리 '희망의🤧 스위치'😇👌🎍😍를 올리자. 찰칵! "

	matches := emoji.FindAll(str)
	totalUniqueEmoji := len(matches)

	assert.Equal(t, totalUniqueEmoji, 8, "There should be one emoji")

	emojiRemoved := emoji.RemoveAll(str)
	assert.Equal(t, "포기하고 싶은 순간들 바로 그 순간 빨리 '희망의 스위치'를 올리자. 찰칵!", emojiRemoved, "There should be no emoji")

}

func TestOutOfRangeError(t *testing.T) {

	str := "武柳💁👌🎍😍昊雨"

	matches := emoji.FindAll(str)
	totalUniqueEmoji := len(matches)

	assert.Equal(t, totalUniqueEmoji, 4, "There should be one emoji")

	emojiRemoved := emoji.RemoveAll(str)
	assert.Equal(t, "武柳昊雨", emojiRemoved, "There should be no emoji")

}
