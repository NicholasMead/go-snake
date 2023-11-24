package ticker

import (
	"testing"
	"time"
)

func TestStop(t *testing.T) {
	exp_duration := 100 * time.Millisecond
	ticker := NewTicker()
	defer ticker.Close()

	ticker.Start()
	time.Sleep(10 * time.Millisecond)

	act_duration := ticker.Stop()

	if act_duration == 0 {
		t.Errorf("Did not stop in time. Got %v", act_duration)
	}
	if act_duration >= FrameLength(exp_duration) {
		t.Errorf("Did not stop in time. Got %v expected < %v", act_duration, exp_duration)
	}
}
