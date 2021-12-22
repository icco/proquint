package proquint

import (
	"strconv"
	"testing"
)

func TestGenerate(t *testing.T) {
	for i := 1; i <= 10; i++ {
		tc := i
		t.Run(strconv.Itoa(tc), func(t *testing.T) {
			t.Parallel()
			v, err := Generate(tc)
			if err != nil {
				t.Error(err)
			}

			if t.Verbose() {
				t.Logf("generated %q for %d", v, tc)
			}

			expectedLen := tc*6 - 1
			if len(v) != expectedLen {
				t.Errorf("expected %d, got %d", expectedLen, len(v))
			}
		})
	}
}

func BenchmarkGenerateThree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tc := 3
		v, err := Generate(tc)
		if err != nil {
			t.Error(err)
		}

		if t.Verbose() {
			t.Logf("generated %q for %d", v, tc)
		}

	}
}
