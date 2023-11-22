package unitspace

type Vector [3]int

func Add(a, b Vector) Vector {
	c := Vector{}
	for i := range c {
		c[i] = a[i] + b[i]
	}
	return c
}

func Distance(a, b Vector) int {
	value := 0
	for i := range a {
		value += Abs(a[i] - b[i])
	}
	return value
}

func (in Vector) forEach(action func(int, int) int) (out Vector) {
	for i, v := range in {
		out[i] = action(i, v)
	}
	return
}
