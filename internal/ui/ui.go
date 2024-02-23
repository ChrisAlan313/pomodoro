package ui

import (
	"context" // Importing the context package for managing request-scoped values, cancelation signals, and deadlines
	"fmt"     // Importing the fmt package for formatted I/O operations
	"time"    // Importing the time package for time-related operations
)

// Run initializes the application's main logic with a specified context, time, and cancel function.
func Run(ctx context.Context, t time.Time) {

	ticker := time.NewTicker(time.Second) // Creates a new ticker that ticks every 500 milliseconds
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
			updateTime(tl) // Calls updateTime with the time left until the deadline
		}
	}
}

// updateTime prints the remaining time until a specified deadline.
func updateTime(tl time.Duration) {
	tl = tl.Truncate(time.Second)
	st := tl.String()
	fmt.Printf("\033[A\033[2K")
	fmt.Println("\rTime left: ", st) // Prints the remaining duration until the deadline
}
