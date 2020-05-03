package b

import (
	"strings"
)

// B replace the letter B and O in a string with the emojis 🅱️ and 🅾️, respectively.
func B(text string) string {
	result := strings.ReplaceAll(text, "b", "🅱️")
	result = strings.ReplaceAll(result, "B", "🅱️")
	return result
}
