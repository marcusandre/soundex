package soundex

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	pairs := map[string]string{
		"Wikipedia":  "W213",
		"Britney":    "B635",
		"Spears":     "S162",
		"SuperZicke": "S162",
		"SMITS":      "S532",
		"Weiß":       "W200",
		"Rupert":     "R163",
		"Robert":     "R163",
	}
	for word, target := range pairs {
		soundex := New(word)
		fmt.Println(soundex)
		if soundex != target {
			t.Errorf("%s does not equal %s", soundex, target)
		}
	}
}
