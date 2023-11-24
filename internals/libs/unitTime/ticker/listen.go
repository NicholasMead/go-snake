package ticker

type listenRequest struct {
	resp chan<- tickResponce
}

func (t *ticker) Listen() (<-chan FrameLength, error) {
	if err := t.errorIfClosed(); err != nil {
		return nil, err
	}
	respc := make(chan tickResponce)

	t.listen <- listenRequest{respc}

	resp := <-respc
	return resp.tick, resp.err
}
