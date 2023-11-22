package snake

import (
	"testing"

	"nickmead.tech/snake/helpers/assert"
	unit "nickmead.tech/snake/internals/libs/unitspace"
)

func TestSnake(t *testing.T) {
	t.Run("GetAll", func(t *testing.T) {
		head := unit.Vector{}
		tail := []unit.Vector{
			{1, 0},
			{2, 0},
			{3, 0},
		}

		s := Snake{head, tail, unit.Vector{}}.Position()

		assert.AssertEqual(4, len(s), t)
		assert.AssertEqual(head, s[0], t)
		for i, v := range tail {
			assert.AssertEqual(v, s[i+1], t)
		}
	})

	t.Run("Turn/prevents_turn_back", func(t *testing.T) {
		snakes := []struct {
			name  string
			snake Snake
		}{
			{
				"normal",
				Snake{
					head:     unit.Vector{1, 0},
					tail:     []unit.Vector{{0, 0}},
					velocity: unit.Vector{},
				},
			},
			{
				"just_eaten",
				Snake{
					head:     unit.Vector{1, 0},
					tail:     []unit.Vector{{1, 0, 1}, {0, 0}},
					velocity: unit.Vector{},
				},
			},
			{
				"eaten_plus_one",
				Snake{
					head:     unit.Vector{1, 0},
					tail:     []unit.Vector{{0, 0, 1}, {0, 0}},
					velocity: unit.Vector{},
				},
			},
		}

		for _, test := range snakes {
			t.Run(test.name, func(t *testing.T) {
				if err := test.snake.Turn(unit.Vector{-1}); err == nil {
					t.Error("Turnback not prevented")
				}
			})
		}
	})

	t.Run("Move", func(t *testing.T) {
		snake := Build(unit.Vector{3, 0}, unit.Vector{0, 0})
		expect := Build(unit.Vector{4, 0}, unit.Vector{1, 0}).Position()

		snake.Turn(unit.Vector{1})
		snake.Move()

		assert.AssertEqualSlice(expect, snake.Position(), t)
	})

	t.Run("Eat", func(t *testing.T) {
		snake := Build(unit.Vector{3, 0}, unit.Vector{0, 0})
		expect := []unit.Vector{
			{3, 0},
			{3, 0, 1},
			{2, 0},
			{1, 0},
			{0, 0},
		}

		snake.Eat()

		assert.AssertEqualSlice(expect, snake.Position(), t)
	})
}

func TestAt(t *testing.T) {
	pos := unit.Vector{2, 3, 0}
	s := At(pos).Position()

	assert.AssertEqual(1, len(s), t)
	assert.AssertEqual(pos, s[0], t)
}

func TestBuild(t *testing.T) {
	from := unit.Vector{1, 1}
	to := unit.Vector{2, 3}

	snake := Build(from, to).Position()

	expect := []unit.Vector{
		{1, 1},
		{2, 1},
		{2, 2},
		{2, 3},
	}
	assert.AssertEqualSlice(expect, snake, t)
}

// func assert.AssertEqual[T comparable](e, r T, t *testing.T) {
// 	if r != e {
// 		t.Errorf("Expected %v got %v}", e, r)
// 	}
// }

// func assert.assert.AssertEqualSlice[T comparable](expect, result []T, t *testing.T) {
// 	if len(expect) != len(result) {
// 		t.Fatalf("Expected len %v got %v", len(expect), len(result))
// 	}
// 	for i, e := range expect {
// 		assert.AssertEqual(e, result[i], t)
// 	}
// }
