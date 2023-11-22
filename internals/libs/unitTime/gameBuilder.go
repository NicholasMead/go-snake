package unitTime

import (
	"nickmead.tech/snake/internals/libs/unitTime/ticker"
)

type GameBuilder[TState any] interface {
	SetInitialState(TState) GameBuilder[TState]
	SetIntitialFrameRate(fps float32) GameBuilder[TState]
	AddController(GameController[TState]) GameBuilder[TState]
	// controllers []GameController[TState]
}

func Build[TState any](delegate func(builder GameBuilder[TState])) Game[TState] {
	builder := gameBuilder[TState]{
		game: Game[TState]{
			ticker:    ticker.NewTicker(),
			frameRate: 1,
		},
	}

	delegate(&builder)

	return builder.game
}

type gameBuilder[TState any] struct {
	game Game[TState]
}

func (b *gameBuilder[TState]) AddController(controller GameController[TState]) GameBuilder[TState] {
	b.game.controllers = append(b.game.controllers, controller)
	return b
}

func (b *gameBuilder[TState]) SetInitialState(data TState) GameBuilder[TState] {
	b.game.state = data
	return b
}

func (builder *gameBuilder[TState]) SetIntitialFrameRate(fps float32) GameBuilder[TState] {
	builder.game.frameRate = fps
	return builder
}
