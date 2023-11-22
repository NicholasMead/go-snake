package unitTime

import "nickmead.tech/snake/internals/libs/unitTime/ticker"

type Game[TState any] struct {
	state       TState
	controllers []GameController[TState]

	ticker    ticker.Ticker
	frameRate float32
}

func (game *Game[TArgs]) framePreTick() {
	for _, c := range game.controllers {
		c.OnFramePreTick(GameContext[TArgs]{game})
	}
}

func (game *Game[TArgs]) framePostTick() {
	for _, c := range game.controllers {
		c.OnFramePostTick(GameContext[TArgs]{game})
	}
}

func (game *Game[TArgs]) Tick() {
	game.framePreTick()
	for _, c := range game.controllers {
		c.OnFrameTick(GameContext[TArgs]{game})
	}
	game.framePostTick()
}
