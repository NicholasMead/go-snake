package ticker

import (
	"errors"
	"time"
)

type Tick time.Duration

type Ticker interface {
	// Performs at tick
	Tick(time.Duration) (<-chan Tick, error)

	// Gets the current in-proress tick or returns an error
	Listen() (<-chan Tick, error)

	// Stops the current in-process tick and return the elapsed time
	Stop() (Tick, error)

	// Closes the Ticker
	Close() (Tick, error)
}

type ticker struct {
	tick   chan tickRequest
	listen chan listenRequest
	stop   chan stopRequest
	close  chan stopRequest
}

func (t *ticker) errorIfClosed() error {
	if t.close == nil || t.tick == nil {
		return errors.New("ticker close")
	}
	return nil
}

func NewTicker() Ticker {
	t := &ticker{
		tick:   make(chan tickRequest),
		listen: make(chan listenRequest),
		stop:   make(chan stopRequest),
		close:  make(chan stopRequest),
	}

	go actor(t)

	return t
}
