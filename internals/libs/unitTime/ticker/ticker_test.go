package ticker

import (
	"testing"
	"time"
)

func TestTicker(t *testing.T) {

	t.Run("CanTick", func(t *testing.T) {
		ticker := NewTicker()
		exp_duration := 1 * time.Millisecond
		defer ticker.Close()

		tick, err := ticker.Tick(exp_duration)

		if err != nil {
			t.Fatal(err)
		}
		if tick == nil {
			t.Fatal("No tick channel returned")
		}
		if act_duration := <-tick; act_duration != FrameLength(exp_duration) {
			t.Errorf("Durations does not match. expected %v for %v", exp_duration, act_duration)
		}
		select {
		case _, ok := <-tick:
			if ok {
				t.Error("Channel not closed")
			}
		default:
			t.Error("Channel not closed")
		}
	})

	t.Run("CannotTickTwice", func(t *testing.T) {
		ticker := NewTicker()
		defer ticker.Close()

		tick, err := ticker.Tick(0)
		if err != nil {
			t.Fatal("First tick failed")
		}

		if _, err := ticker.Tick(0); err == nil {
			t.Fatal("Second tick didn't fail")
		}

		<-tick

		if _, err := ticker.Tick(0); err != nil {
			t.Fatal("Second tick didn't fail")
		}

	})

	t.Run("CanClose", func(t *testing.T) {
		ticker := NewTicker()

		ticker.Tick(1 * time.Second)
		tick, err := ticker.Close()

		if err != nil {
			t.Error("Close Failed")
		}
		if tick >= FrameLength(1*time.Second) {
			t.Error("Did not close in time")
		}
		if _, err := ticker.Tick(0); err == nil {
			t.Error("Did not close")
		}
	})

	t.Run("CanListen", func(t *testing.T) {
		ticker := NewTicker()
		duration := 10 * time.Millisecond
		defer ticker.Close()

		if _, err := ticker.Tick(duration); err != nil {
			t.Error("Failed first tick")
		}

		tick, err := ticker.Listen()

		if err != nil {
			t.Fatalf("Could not start listen. %v", err)
		}
		if _t := <-tick; _t != FrameLength(duration) {
			t.Error("Did not get correct tick")
		}
	})
}
