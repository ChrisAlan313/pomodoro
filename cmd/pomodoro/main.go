package main

import (
	"context"
	"github.com/chrisalan313/pomodoro/internal/ui"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ui.Run(ctx)
}
