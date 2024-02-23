package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chrisalan313/pomodoro/internal/ui"
)

func main() {
	fmt.Println("Entered main")
	defer fmt.Println("Exited main")
	ctx, canc := context.WithCancel(context.Background())
	defer canc()

	// Amount of time we have for this pomodoro
	// Make this something we can pass in later
	t := 5 * time.Second

	ui.Run(ctx, time.Now().Add(t))
}
