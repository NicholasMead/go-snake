package unitspace

import (
	"fmt"
	"testing"

	"nickmead.tech/snake/helpers/assert"
)

func TestAbs(t *testing.T) {
	tests := []struct{ in, out int }{
		{+0, 0},
		{+1, 1},
		{-1, 1},
		{+3, 3},
		{-3, 3},
	}

	for _, test := range tests {
		r := Abs(test.in)

		assert.AssertEqual[int](test.out, r, t)
	}
}

func BenchmarkAbs(b *testing.B) {
	b.Run("abs_impl_direct", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			abs_impl_direct(i)
		}
	})
	b.Run("abs_impl_math", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			abs_impl_math(i)
		}
	})
}

func TestNormalise(t *testing.T) {
	cases := []struct{ in, out int }{
		{+0, +0},
		{+1, +1},
		{-1, -1},
		{+2, +1},
		{-2, -1},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.in), func(t *testing.T) {
			assert.AssertEqual(c.out, Normalise(c.in), t)
		})
	}
}
