package ticker

func (t *ticker) Close() (Tick, error) {
	if err := t.errorIfClosed(); err != nil {
		return 0, err
	}
	resp := make(chan Tick)

	t.close <- resp

	*t = ticker{}

	return <-resp, nil
}
