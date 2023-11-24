package ticker

type stopRequest chan<- FrameLength

func (t *ticker) Stop() (FrameLength, error) {
	if err := t.errorIfClosed(); err != nil {
		return 0, err
	}

	resp := make(chan FrameLength)
	req := resp

	t.stop <- req

	return <-req, nil
}
