// proquint is a simplified implementation of https://arxiv.org/html/0901.4016
// to make UUIDs that are easy to pronounce.
package proquint

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	consonants = []string{"b", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "s", "t", "v", "w", "z"}
	vowels     = []string{"a", "e", "i", "o", "u", "r"}
)

// Generate creates a proquint of of the length of words, separated by a dash.
func Generate(words int) (string, error) {
	if words <= 0 {
		return "", fmt.Errorf("cannot generate less than one word")
	}

	rand.Seed(time.Now().Unix())

	var ret []string

	for i := 0; i < words; i++ {
		var word []string
		for j := 0; j < 5; j++ {
			var letter string
			if j%2 != 0 {
				letter = vowels[rand.Intn(len(vowels))]
			} else {
				letter = consonants[rand.Intn(len(consonants))]
			}

			word = append(word, letter)
		}
		ret = append(ret, strings.Join(word, ""))
	}

	return strings.Join(ret, "-"), nil
}
