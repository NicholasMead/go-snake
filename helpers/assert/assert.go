package assert

import "testing"

func AssertEqual[T comparable](e, r T, t *testing.T) {
	if r != e {
		t.Errorf("Expected %v got %v}", e, r)
	}
}

func AssertEqualSlice[T comparable](expect, result []T, t *testing.T) {
	if len(expect) != len(result) {
		t.Fatalf("Expected len %v got %v", len(expect), len(result))
	}
	for i, e := range expect {
		AssertEqual(e, result[i], t)
	}
}
