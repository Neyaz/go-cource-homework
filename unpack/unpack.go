package unpack

import (
	"unicode"
	"strings"
)

// Unpack recieves string and unpuck it
func Unpack(str string) string {
	var prev rune
	var b strings.Builder
	for _, rune := range str {
		if unicode.IsDigit(rune) {
			if unicode.IsDigit(prev) {
				return ""
			}
			for i := 1; i <= (int(rune) - '0')-1; i++ {
					b.WriteString(string(prev))
			}
		} else {
			b.WriteString(string(rune))
		}
		prev = rune
	}

	return b.String()
}