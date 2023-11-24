package unitTime

import "sync"

type GameController[TState any] interface {
	OnInputRecieved(GameContext[TState])

	OnPreFrame(GameContext[TState])
	OnFrame(GameContext[TState])
	OnPostFrame(GameContext[TState])
}

type BaseController[TState any] struct{}

func (BaseController[TState]) OnInputRecieved(ctx GameContext[TState]) {
}

func (BaseController[TState]) OnPreFrame(ctx GameContext[TState]) {
}

func (BaseController[TState]) OnFrame(ctx GameContext[TState]) {
}

func (BaseController[TState]) OnPostFrame(ctx GameContext[TState]) {
}

type gameContollerGroup[TState any] struct {
	controllers []GameController[TState]
}

func (g gameContollerGroup[TState]) doAction(action func(c GameController[TState])) {
	var wg sync.WaitGroup

	for _, c := range g.controllers {
		wg.Add(1)
		go func(c GameController[TState]) {
			defer wg.Done()
			action(c)
		}(c)
	}

	wg.Wait()
}

func (g gameContollerGroup[TState]) OnInputRecieved(ctx GameContext[TState]) {
	g.doAction(func(c GameController[TState]) { c.OnInputRecieved(ctx) })
}

func (g gameContollerGroup[TState]) OnPreFrame(ctx GameContext[TState]) {
	g.doAction(func(c GameController[TState]) { c.OnPreFrame(ctx) })
}

func (g gameContollerGroup[TState]) OnFrame(ctx GameContext[TState]) {
	g.doAction(func(c GameController[TState]) { c.OnFrame(ctx) })
}

func (g gameContollerGroup[TState]) OnPostFrame(ctx GameContext[TState]) {
	g.doAction(func(c GameController[TState]) { c.OnPostFrame(ctx) })
}
