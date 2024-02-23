package ui

import (
	"context" // Importing the context package for managing request-scoped values, cancelation signals, and deadlines
	"fmt"     // Importing the fmt package for formatted I/O operations
	"time"    // Importing the time package for time-related operations
)

// Run initializes the application's main logic with a specified context, time, and cancel function.
func Run(ctx context.Context, t time.Time) {
	fmt.Println("Entered Run")      // Prints a message indicating the start of the Run function
	defer fmt.Println("Exited Run") // Ensures that a message indicating the exit of the Run function is printed at the end

	ticker := time.NewTicker(500 * time.Millisecond) // Creates a new ticker that ticks every 500 milliseconds
	defer ticker.Stop()                              // Ensures that the ticker is stopped when the function exits

	fmt.Println("Set tick")     // Prints a message indicating that the ticker has been set
	eventLoop(ctx, t, ticker.C) // Starts the eventLoop function in a new goroutine
}

// eventLoop continuously checks for events such as tick events or context cancellation.
func eventLoop(ctx context.Context, dl time.Time, tick <-chan time.Time) {
	fmt.Println("Entered eventLoop")      // Prints a message indicating the start of the eventLoop function
	defer fmt.Println("Exited eventLoop") // Ensures that a message indicating the exit of the eventLoop function is printed at the end

	for {
		select {
		case <-tick:
			fmt.Println("tick happened")
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
	fmt.Println("Entered updateTime")            // Prints a message indicating the start of the updateTime function
	st := tl.String()                            // Converts the duration to a string
	fmt.Println("Duration until deadline: ", st) // Prints the remaining duration until the deadline
	fmt.Println("Exited updateTime")             // Prints a message indicating the exit of the updateTime function
}
