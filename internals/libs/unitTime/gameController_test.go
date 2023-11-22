package unitTime

import "testing"

type data struct{}

func TestBaseController(t *testing.T) {
	t.Run("isGameController", func(t *testing.T) {
		var _ GameController[data] = BaseController[data]{}
	})
}
