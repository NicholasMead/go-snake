package unitspace

import "math"

func Normalise(a int) int {
	if a == 0 {
		return 0
	}
	return a / Abs(a)
}

func Abs(a int) int {
	// goos: windows
	// goarch: amd64
	// pkg: nickmead.tech/snake/internals/domain/space/unitspace
	// cpu: Intel(R) Core(TM) i9-9900KF CPU @ 3.60GHz
	// BenchmarkAbs/abs_impl_direct-16                 1000000000               0.2223 ns/op
	// BenchmarkAbs/abs_impl_math-16                   1000000000               0.2189 ns/op
	return abs_impl_direct(a)
}

func Min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func Max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func abs_impl_direct(a int) int {
	if a == 0 {
		return 0
	} else if a < 0 {
		return -a
	} else {
		return a
	}
}

func abs_impl_math(a int) int {
	return int(math.Abs(float64(a)))
}
