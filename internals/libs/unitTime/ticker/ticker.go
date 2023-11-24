package ticker

import (
	"errors"
	"time"
)

type FrameLength time.Duration
type FrameRate float32

func (t FrameLength) ToFrameRate() FrameRate {
	return FrameRate(float32(time.Second) / float32(t))
}

func (fr FrameRate) ToFrameLength() FrameLength {
	return FrameLength(float32(fr) / float32(time.Second))
}

func Fps(fps float32) FrameRate {
	return FrameRate(fps)
}

type Ticker interface {
	// Gets the current in-proress tick or returns an error
	Listen() <-chan FrameLength

	// Starts the tickers
	Start()

	// Stops the current in-process tick and return the elapsed time
	Stop() FrameLength

	// Sets the Frame Rate
	SetFrameRate(fr FrameRate)

	// Adjusts the Frame Rate by a set factor
	AdjustFrameRate(factor float32)

	// Closes the Ticker
	Close() FrameLength
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
