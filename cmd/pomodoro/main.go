package main

import (
	"context"
	"github.com/chrisalan313/pomodoro/internal/ui"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Amount of time we have for this pomodoro
	// Make this something we can pass in later
	t := 5 * time.Second
	ui.Run(ctx, time.Now().Add(t))
}
