package ticker

func (t *ticker) Close() (FrameLength, error) {
	if err := t.errorIfClosed(); err != nil {
		return 0, err
	}
	resp := make(chan FrameLength)

	t.close <- resp

	*t = ticker{}

	return <-resp, nil
}
