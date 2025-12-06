// proquint is a simplified implementation of https://arxiv.org/html/0901.4016
// to make UUIDs that are easy to pronounce.
package proquint

import (
	"fmt"
	"math/big"
	"crypto/rand"
	"strings"
)

var (
	consonants = []string{"b", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
	vowels     = []string{"a", "i", "o", "u"}
)

// Generate creates a proquint of the length of words, separated by a dash.
//
// https://merveilles.town/@flbr/115574222078642407
func Generate(words int) (string, error) {
	if words <= 0 {
		return "", fmt.Errorf("cannot generate less than one word")
	}

	var ret []string
	max := big.NewInt(2 << 15)
	for i := 0; i < words; i++ {
		bigN, err := rand.Int(rand.Reader, max)
		n := bigN.Int64()
		if err != nil {
			return "", err
		}
		word := fmt.Sprintf("%s%s%s%s%s", consonants[n>>12], vowels[(n>>10)&3], consonants[(n>>6)&15], vowels[(n>>4)&3], consonants[n&15])
		ret = append(ret, word)
	}

	return strings.Join(ret, "-"), nil
}
