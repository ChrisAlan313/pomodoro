package main

import (
	"context"
	"time"

	"github.com/chrisalan313/pomodoro/internal/ui"
)

func main() {
	ctx, canc := context.WithCancel(context.Background())
	defer canc()

	// Amount of time we have for this pomodoro
	// Make this something we can pass in later
	t := 25 * time.Minute

	ui.Run(ctx, time.Now().Add(t))
}
