package unitspace

import (
	"testing"

	"nickmead.tech/snake/helpers/assert"
)

func TestAdd(t *testing.T) {
	v1 := Vector{1, 2, 3}
	v2 := Vector{2, 4, 6}

	r := Add(v1, v2)
	e := Vector{3, 6, 9}

	assert.AssertEqual(e, r, t)
}

func TestMeasure(t *testing.T) {
	v1 := Vector{1, 1, 1}
	v2 := Vector{2, 2, 2}

	r := []int{Distance(v1, v2), Distance(v2, v1)}
	e := 3

	for _, v := range r {
		assert.AssertEqual(e, v, t)
	}
}
