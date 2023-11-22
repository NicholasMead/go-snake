package unitspace

import "math"

type Bound [2]Vector

func (bound Bound) Wrap(in Vector) (out Vector) {
	// bound can describe any two corners of a rectange, so lets pars the bottom left and top right
	bound = bound.normalise()

	for i, v := range in {
		size := bound[1][i] - bound[0][i] + 1

		if size != 0 {
			out[i] = v - bound[0][i]
			out[i] %= size
			out[i] += bound[0][i]
		}
	}

	return out
}

func (b Bound) normalise() Bound {
	start := Vector{math.MaxInt, math.MaxInt}
	end := Vector{math.MinInt, math.MinInt}

	for _, vector := range b {
		for i, value := range vector {
			start[i] = Min(start[i], value)
			end[i] = Max(end[i], value)
		}
	}

	return Bound{start, end}
}
