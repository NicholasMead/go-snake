package unitTime

import "nickmead.tech/snake/internals/libs/unitTime/ticker"

type Game[TState any] struct {
	state      TState
	controller GameController[TState]

	ticker ticker.Ticker
}

func (game *Game[TArgs]) onFrame() {
	game.controller.OnPreFrame(GameContext[TArgs]{game})
	game.controller.OnFrame(GameContext[TArgs]{game})
	game.controller.OnPostFrame(GameContext[TArgs]{game})
}
