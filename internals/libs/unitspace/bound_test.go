package unitspace

import (
	"fmt"
	"testing"

	"nickmead.tech/snake/helpers/assert"
)

func TestBound(t *testing.T) {
	t.Run("Wrap", func(t *testing.T) {

		test := []struct {
			bound Bound
			in    Vector
			out   Vector
		}{
			{Bound{Vector{4, 4}, Vector{0, 0}}, Vector{6, 6}, Vector{1, 1}},
			{Bound{Vector{0, 0}, Vector{4, 4}}, Vector{6, 6}, Vector{1, 1}},
			{Bound{Vector{0, 4}, Vector{4, 0}}, Vector{7, 7}, Vector{2, 2}},
			{Bound{Vector{4, 0}, Vector{0, 4}}, Vector{7, 7}, Vector{2, 2}},
			{Bound{Vector{1, 1}, Vector{5, 5}}, Vector{6, 6}, Vector{1, 1}},
			{Bound{Vector{1, 1}, Vector{4, 4}}, Vector{6, 6}, Vector{2, 2}},
		}

		for i := range test {
			t.Run(fmt.Sprintf("%v", test[i]), func(t *testing.T) {
				wrapped := test[i].bound.Wrap(test[i].in)

				assert.AssertEqual(test[i].out, wrapped, t)
			})
		}
	})
}
