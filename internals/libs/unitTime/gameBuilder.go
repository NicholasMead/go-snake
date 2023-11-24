package unitTime

import "nickmead.tech/snake/internals/libs/unitTime/ticker"

type GameBuilder[TState any] interface {
	SetInitialState(TState) GameBuilder[TState]
	SetIntitialFrameRate(fps float32) GameBuilder[TState]
	AddController(GameController[TState]) GameBuilder[TState]
}

func Build[TState any](delegate func(builder GameBuilder[TState])) Game[TState] {
	builder := gameBuilder[TState]{
		frameRate: 1,
	}

	delegate(&builder)

	game := Game[TState]{
		state:      builder.state,
		controller: builder.controllers,

		ticker: ticker.NewTicker(),
	}
	return game
}

type gameBuilder[TState any] struct {
	state       TState
	controllers gameContollerGroup[TState]

	frameRate float32
}

func (b *gameBuilder[TState]) AddController(controller GameController[TState]) GameBuilder[TState] {
	b.controllers = gameContollerGroup[TState]{
		append(b.controllers.controllers, controller),
	}
	return b
}

func (b *gameBuilder[TState]) SetInitialState(data TState) GameBuilder[TState] {
	b.state = data
	return b
}

func (builder *gameBuilder[TState]) SetIntitialFrameRate(fps float32) GameBuilder[TState] {
	builder.frameRate = fps
	return builder
}
