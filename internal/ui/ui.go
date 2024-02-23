package ui

import (
	"fmt"  // Importing the fmt package for formatted I/O operations
	"time" // Importing the time package for time-related operations
)

// updateTime prints the remaining time until a specified deadline.
func UpdateTime(tl time.Duration) {
	tl = tl.Truncate(time.Second)
	st := tl.String()
	fmt.Printf("\033[A\033[2K")
	fmt.Println("\rTime left: ", st) // Prints the remaining duration until the deadline
}
