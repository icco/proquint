package proquint

import (
	"strconv"
	"testing"
)

func TestGenerateFail(t *testing.T) {
	tests := []int{-1, 0}
	for _, i := range tests {
		tc := i
		t.Run(strconv.Itoa(tc), func(t *testing.T) {
			t.Parallel()
			_, err := Generate(tc)
			if err == nil {
				t.Errorf("expected error for %d", tc)
			}
		})
	}
}

func TestGenerate(t *testing.T) {
	for i := 1; i <= 1000; i++ {
		tc := i
		t.Run(strconv.Itoa(tc), func(t *testing.T) {
			t.Parallel()
			v, err := Generate(tc)
			if err != nil {
				t.Error(err)
			}

			if testing.Verbose() {
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
			b.Error(err)
		}

		if testing.Verbose() {
			b.Logf("generated %q for %d", v, tc)
		}

	}
}
