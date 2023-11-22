package ticker

import "time"

type tickRequest struct {
	duration time.Duration
	resp     chan<- tickResponce
}

type tickResponce struct {
	tick <-chan Tick
	err  error
}

func (t *ticker) Tick(duration time.Duration) (<-chan Tick, error) {
	if err := t.errorIfClosed(); err != nil {
		return nil, err
	}
	respc := make(chan tickResponce)

	t.tick <- tickRequest{duration, respc}

	resp := <-respc
	return resp.tick, resp.err
}
