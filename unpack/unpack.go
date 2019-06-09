package unpack

import (
	"unicode"
)

// Unpack recieves string and unpuck it
func Unpack(str string) string {
	var result string
	var prev rune
	for _, rune := range str {
		if unicode.IsDigit(rune) {
			if unicode.IsDigit(prev) {
				return ""
			}
			for i := 1; i <= (int(rune) - '0')-1; i++ {
					result = result + string(prev)
			}
		} else {
			result = result + string(rune)
		}
		prev = rune
	}

	return result
}