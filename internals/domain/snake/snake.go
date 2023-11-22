package snake

import (
	"errors"

	unit "nickmead.tech/snake/internals/libs/unitspace"
)

type Snake struct {
	head     unit.Vector
	tail     []unit.Vector
	velocity unit.Vector
}

func At(pos unit.Vector) *Snake {
	return &Snake{pos, []unit.Vector{}, unit.Vector{}}
}

func Build(from, to unit.Vector) *Snake {
	snake := Snake{from, []unit.Vector{}, unit.Vector{}}
	current := from

	for dim := range (unit.Vector{}) {
		for current[dim] != to[dim] {
			diff := unit.Normalise(to[dim] - current[dim])
			step := unit.Vector{}
			step[dim] = unit.Normalise(diff)
			current = unit.Add(current, step)

			snake.tail = append(snake.tail, current)
		}
	}

	return &snake
}

func (s Snake) Position() []unit.Vector {
	body := []unit.Vector{s.head}
	body = append(body, s.tail...)

	return body
}

func (s Snake) Velocity() unit.Vector {
	return s.velocity
}

func (s *Snake) Turn(velocity unit.Vector) error {
	//check first two tail pieces for conflict
	for i := 0; i < len(s.tail) && i < 2; i++ {
		if unit.Add(s.head, velocity) == s.tail[i] {
			return errors.New("turnback not allowed")
		}
	}

	s.velocity = velocity
	return nil
}

func (s *Snake) Move() {
	s.tail = append([]unit.Vector{s.head}, s.tail[:len(s.tail)-1]...)
	s.head = unit.Add(s.head, s.velocity)
}

func (s *Snake) Eat() {
	next := unit.Add(s.head, unit.Vector{0, 0, 1})
	s.tail = append([]unit.Vector{next}, s.tail...)
}

func (s *Snake) Wrap(b unit.Bound) {
	s.head = b.Wrap(s.head)
	for i, v := range s.tail {
		s.tail[i] = b.Wrap(v)
	}
}
