package arena

import (
	"testing"

	unit "nickmead.tech/snake/internals/libs/unitspace"
)

type stubItem struct {
	position unit.Vector
}

func (i *stubItem) Compose() []Block {
	return []Block{{"Item", i.position}}
}

// func (i *stubItem) Wrap(b unit.Bound) {
// 	i.position = b.Wrap(i.position)
// }

func TestArena(t *testing.T) {
	t.Run("Items", func(t *testing.T) {
		a := Arena{}
		var item = &stubItem{}

		a.AddItem(item)
		result := a.GetItems()[0]

		item.position = unit.Vector{1, 1}
		if item != result {
			t.Errorf("Items did not match: %v, %v", item, result)
		}
	})

	t.Run("Compose", func(t *testing.T) {
		item := &stubItem{}
		a := Arena{
			unit.Bound{},
			[]Item{item},
		}

		blocks := a.Compose()

		if len(blocks) != 1 {
			t.Fatalf("Expected 1 block, got %v", len(blocks))
		}
		if blocks[0] != item.Compose()[0] {
			t.Fatalf("Expected %v block, got %v", item.Compose()[0], blocks[0])
		}
	})
}
