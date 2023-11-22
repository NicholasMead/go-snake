package domain

import (
	"nickmead.tech/snake/internals/domain/arena"
	"nickmead.tech/snake/internals/domain/snake"
	"nickmead.tech/snake/internals/libs/unitTime"
)

type State struct {
	arena arena.Arena
	snake *snake.Snake
}

type GameContext = unitTime.GameContext[State]

type Game = unitTime.Game[State]

type SnakeController struct {
	unitTime.BaseController[State]
}

func (SnakeController) OnFrameTick(ctx GameContext) {
	ctx.State().snake.Move()
}

func (SnakeController) OnFrameProstTick(ctx GameContext) {
	data := ctx.State()
	snake := data.snake
	arena := data.arena

	snake.Wrap(arena.Bound)
}

func add() {
	var _ = unitTime.Build[State](func(builder unitTime.GameBuilder[State]) {
		builder.AddController(SnakeController{})
	})
}

// Snake() *snake.Snake

// AddPoints(int)

// type GameController interface {
// 	arena.Composable

// 	PerformPreTickActions(GameContext)
// 	PerformOnTickActions(GameContext)
// 	PerformPostTickAction(GameContext)
// }
