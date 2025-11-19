// proquint is a simplified implementation of https://arxiv.org/html/0901.4016
// to make UUIDs that are easy to pronounce.
package proquint

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

var (
	consonants = []string{"b", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "s", "t", "v", "w", "z"}
	vowels     = []string{"a", "e", "i", "o", "u", "r"}
)

// Generate creates a proquint of of the length of words, separated by a dash.
//
// https://merveilles.town/@flbr/115574222078642407
func Generate(words int) (string, error) {
	if words <= 0 {
		return "", fmt.Errorf("cannot generate less than one word")
	}

	var ret []string

	for i := 0; i < words; i++ {
		n := rand.Intn(int(math.Pow(2, 16)))
		word := fmt.Sprintf("%s%s%s%s%s", consonants[n>>12], vowels[(n>>10)&3], consonants[(n>>6)&15], vowels[(n>>4)&3], consonants[n&15])
		ret = append(ret, word)
	}

	return strings.Join(ret, "-"), nil
}
