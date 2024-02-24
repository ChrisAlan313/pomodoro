package main

import (
	"context"
	"time"

	"github.com/chrisalan313/pomodoro/internal/audio"
	"github.com/chrisalan313/pomodoro/internal/ui"
)

func main() {
	ctx, canc := context.WithCancel(context.Background())
	defer canc()

	// Get music started
	go audio.Play("/Users/cbooty/Music/underwater JUNGLE DNB   a jungle drum and bass mix.mp3")

	// Amount of time we have for this pomodoro
	// Make this something we can pass in later
	d := 25 * time.Second
	t := time.Now().Add(d)

	ticker := time.NewTicker(time.Second) // Creates a new ticker that ticks every second
	defer ticker.Stop()                   // Ensures that the ticker is stopped when the function exits

	eventLoop(ctx, t, ticker.C) // Starts the eventLoop function in a new goroutine
}

// eventLoop continuously checks for events such as tick events or context cancellation.
func eventLoop(ctx context.Context, dl time.Time, tick <-chan time.Time) {

	for {
		select {
		case <-tick:
			tl := time.Until(dl) // Calculates the duration until the specified deadline
			if tl < 0 {          // Checks if the deadline has passed
				return
			}
			ui.UpdateTime(tl) // Calls updateTime with the time left until the deadline
		}
	}
}
