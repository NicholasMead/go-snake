package arena

import "nickmead.tech/snake/internals/libs/unitspace"

type Arena struct {
	unitspace.Bound
	items []Item
}

func (a *Arena) AddItem(item Item) {
	a.items = append(a.items, item)
}

func (a Arena) GetItems() []Item {
	return a.items
}

func (a Arena) Compose() []Block {
	blocks := []Block{}
	for _, item := range a.items {
		blocks = append(blocks, item.Compose()...)
	}
	return blocks
}

type Composable interface {
	Compose() []Block
}

type Boundable interface {
	Wrap(Arena)
}

type Item interface {
	Composable
	// Boundable
}

type Block struct {
	Name     string
	Position unitspace.Vector
}
