package ui

import (
	"context"
	"time"

	"github.com/rivo/tview"
)

var (
	view *tview.Modal
	app  *tview.Application
)

func Run(ctx context.Context, t time.Time) {
	app = tview.NewApplication()
	view = tview.NewModal().
		SetText("REPLACE THIS WITH TIME COUNT DOWN").
		AddButtons([]string{"Quit", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Quit" {
				app.Stop()
			}
		})
	tick := time.Tick(500 * time.Millisecond)

	go eventLoop(ctx, t, tick)
}

func eventLoop(ctx context.Context, dl time.Time, tick <-chan time.Time) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-tick:
			updateTime(dl)
			// call func that updates UI to reflect new time
		}
	}
}

func updateTime(dl time.Time) {
	// func that updates UI with new string
}

// func that makes new string given time left
