// Package soundex creates a phonetic soundex key for a given word
// If you want to learn more, you can look it up in Wikipedia:
// https://en.wikipedia.org/wiki/Soundex
//
// Side notes:
// - vowels are ignored completely if they are not the first letter
// - same counts for the consonants H, W and Y
// - duplicate codes inside the resulting strings will be removed
// - this is very focused on the English language
// - and beware: it provides only a very rough analysis
package soundex

import (
	"unicode"
)

// MinLength indicates the default minimum length of a soundex
const MinLength = 4

// Characters matched to their soundex-code
var codes = map[rune]rune{
	'b': '1',
	'f': '1',
	'p': '1',
	'v': '1',
	'c': '2',
	'g': '2',
	'j': '2',
	'k': '2',
	'q': '2',
	's': '2',
	'ß': '2',
	'x': '2',
	'z': '2',
	'd': '3',
	't': '3',
	'l': '4',
	'm': '5',
	'n': '5',
	'r': '6',
}

// New creates a soundex for the given string
func New(s string) string {
	var prev rune
	var result []rune

	in := []rune(s)

	for i, r := range in {
		r = unicode.ToLower(r)
		c, codeExists := codes[r]

		switch i {
		case 0:
			result = append(result, unicode.ToUpper(r))
		default:
			if codeExists && prev != c {
				result = append(result, c)
			} else {
				continue
			}
			prev = c
		}
	}

	return padRight(string(result), "0", MinLength)
}

// Adds given char to the right until the given length is reached
func padRight(str, pad string, length int) string {
	for {
		str += pad
		if len(str) > length {
			return str[0:length]
		}
	}
}
