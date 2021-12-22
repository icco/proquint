package proquint

import "testing"

func TestGenerate(t *testing.T) {
	for i := 0; i <= 1000; i++ {
		tc := i // capture range variable
		t.Run(tc, func(t *testing.T) {
			t.Parallel()
			v, err := Generate(tc)
			if err != nil {
				t.Error(err)
			}

			expectedLen := tc*6 - 1
			if len(v) != expectedLen {
				t.Errorf("expected $d, got %d", expectedLen, len(v))
			}
		})
	}
}

func BenchmarkGenerateThree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate(3)
	}
}
