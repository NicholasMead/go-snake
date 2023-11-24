package ticker

import (
	"errors"
	"time"
)

type state struct {
	start time.Time
	end   time.Time
	ready bool
	act   <-chan time.Time
	out   chan FrameLength
}

func (s *state) inProgress() bool {
	return *s != *new(state)
}

func (s state) elapsed() FrameLength {
	return FrameLength(time.Since(s.start))
}

func (s state) duration() FrameLength {
	return FrameLength(s.end.Sub(s.start))
}

func (s *state) close() {
	close(s.out)
}

func (s *state) reset() {
	*s = *new(state)
}
func (s *state) closeAndReset() {
	s.close()
	s.reset()
}

func actor(t *ticker) {

	var (
		// request channels
		req_tick   = t.tick
		req_listen = t.listen
		req_stop   = t.stop
		req_close  = t.close

		// state
		tick state
	)

	for {
		out := tick.out
		if !tick.ready {
			out = nil
		}

		select {
		case req := <-req_tick: // Request for Tick
			if tick.inProgress() {
				req.resp <- tickResponce{nil, errors.New("tick in progress")}
				continue
			}

			start := time.Now()
			tick = state{
				start: start,
				end:   start.Add(req.duration),
				act:   time.After(req.duration),
				out:   make(chan FrameLength),
			}

			req.resp <- tickResponce{tick.out, nil}

		case req := <-req_listen: // Request to listen
			if tick.inProgress() {
				req.resp <- tickResponce{tick.out, nil}
			} else {
				req.resp <- tickResponce{nil, errors.New("no tick in progress")}
			}

		case req := <-req_stop: // Stop the tick
			if tick.inProgress() {
				req <- tick.elapsed()
				tick.closeAndReset()
			} else {
				req <- 0
			}

		case req := <-req_close: // Close the ticker
			if tick.inProgress() {
				req <- tick.elapsed()
				tick.close()
			} else {
				req <- 0
			}
			return

		case <-tick.act: // tick action ready
			tick.ready = true

		case out <- tick.duration(): // Send the tick (none-blocking)
			tick.closeAndReset()
		}
	}
}
