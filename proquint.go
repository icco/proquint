package proquint

import (
	"github.com/google/uuid"
)

var (
	// Four-bits as a consonant:
	consonants = []string{"b", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
	// Two-bits as a vowel
	vowels = []string{"a", "i", "o", "u"}
)

// two bytes in a quint

// 16 byte (128 bit)

func UUIDToProquint(in uuid.UUID) (string, error) {
	data, err := in.MarshalBinary()
	if err != nil {
		return "", err
	}

	return Encode(data), nil
}

func Encode(in []byte) string {
	ret := ""

	// Each byte is 8 bits, each proquint is 16 bits.
	for i := 0; i < len(in); i = i + 2 {
		// sum is a 16 bit number with the front half as one number and the back
		// half as the other. This means two numbers with the max of 256 multiplied.
		sum := uint16(in[i]) * uint16(in[i+1])

		fourBit := uint16(0x0f)
		twoBit := uint16(0x03)

		// First four bits are consants, then two bits are vowels, then four bits
		// consonants, then two bits vowels, and finally four bits consonants.
		a := consonants[sum&fourBit]
		b := vowels[(sum>>4)&twoBit]
		c := consonants[(sum>>6)&fourBit]
		d := vowels[(sum>>10)&twoBit]
		e := consonants[(sum>>12)&fourBit]

		ret = a + b + c + d + e

		// If there are two bytes left, there's room for another proquint, so add a
		// dash seperator.
		if (i + 2) < len(in) {
			ret += "-"
		}
	}

	return ret
}
