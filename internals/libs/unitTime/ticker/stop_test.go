package ticker

import (
	"testing"
	"time"
)

func TestStop(t *testing.T) {
	ticker := NewTicker()
	defer ticker.Close()

	exp_duration := 100 * time.Millisecond

	if _, err := ticker.Tick(exp_duration); err != nil {
		t.Fatalf("Failed to tick. %v", err)
	}

	time.Sleep(10 * time.Millisecond)
	act_duration, err := ticker.Stop()

	if err != nil {
		t.Fatalf("Failed to stop. %v", err)
	}
	if act_duration == 0 {
		t.Errorf("Did not stop in time. Got %v", act_duration)
	}
	if act_duration >= Tick(exp_duration) {
		t.Errorf("Did not stop in time. Got %v expected < %v", act_duration, exp_duration)
	}
}
