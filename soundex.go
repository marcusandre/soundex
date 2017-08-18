// Package soundex creates a phonetic soundex key for a given word.
// If you want to learn more, you can look it up in Wikipedia:
// https://en.wikipedia.org/wiki/Soundex
//
// Some side notes: Vowels are ignored completely if they are not the first letter.
// Same counts for the consonants H, W and Y.
// Duplicate codes inside the resulting strings will be removed.
// This is also very focused on the English language.
// And beware: it provides only a very rough analysis.
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

// Soundex is the phonetic key for a given word
type Soundex struct {
	word string
	key  string
}

// New initializes a new Soundex
func New(word string) Soundex {
	key := generate(word)
	return Soundex{
		word: word,
		key:  key,
	}
}

// String prints the resulting Soundex key
func (s Soundex) String() string {
	return s.key
}

// Generate creates a new soundex from a given word
func generate(word string) string {
	var prev rune
	var result []rune

	in := []rune(word)

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
