package clap

import (
	"strings"
)

// Clap adds clap emoji between words or letters
func Clap(text string) string {
	return Emoji(text, "👏️")
}

// Emoji adds an emoji between words or letters, when text is a sentence or when it's a word, respectively
func Emoji(text string, emoji string) string {
	// If text is a sentence (with spaces)
	if strings.Contains(text, " ") {
		return BetweenWords(text, emoji)
	}
	// If text is a word (no space)
	return BetweenLetters(text, emoji)
}

// BetweenWords adds emoji between words by replacing the spaces with the emoji
func BetweenWords(sentence string, emoji string) string {
	return strings.ReplaceAll(sentence, " ", emoji)
}

// BetweenLetters adds emoji between words by replacing the spaces with the emoji
func BetweenLetters(sentence string, emoji string) string {
	return strings.ToUpper(strings.Join(strings.Split(sentence, ""), emoji))
}
