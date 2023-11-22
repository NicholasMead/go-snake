package unitTime

type GameController[TState any] interface {
	OnInput(GameContext[TState])

	OnFramePreTick(GameContext[TState])
	OnFrameTick(GameContext[TState])
	OnFramePostTick(GameContext[TState])
}

type BaseController[TState any] struct{}

func (BaseController[TState]) OnInput(ctx GameContext[TState]) {
}

func (BaseController[TState]) OnFramePreTick(ctx GameContext[TState]) {
}

func (BaseController[TState]) OnFrameTick(ctx GameContext[TState]) {
}

func (BaseController[TState]) OnFramePostTick(ctx GameContext[TState]) {
}
