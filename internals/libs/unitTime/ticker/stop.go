package ticker

type stopRequest chan<- Tick

func (t *ticker) Stop() (Tick, error) {
	if err := t.errorIfClosed(); err != nil {
		return 0, err
	}

	resp := make(chan Tick)
	req := resp

	t.stop <- req

	return <-req, nil
}
