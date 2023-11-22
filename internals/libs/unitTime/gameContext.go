package unitTime

type GameContext[TArgs any] struct {
	game *Game[TArgs]
}

func (ctx *GameContext[TState]) State() TState {
	return ctx.game.state
}
