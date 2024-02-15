package ui

import (
	"log"
	"strconv"
	"time"

	tui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func Run() {
	if err := tui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer tui.Close()

	p := newCounterBox()
	tui.Render(p)

	for i := 10; i >= 0; i-- {
		p.Text = "Countdown: " + strconv.Itoa(i) + "s"
		tui.Render(p)
		time.Sleep(1 * time.Second)
	}
}

func rectDimensions(width, height int) (int, int, int, int) {
	termWidth, termHeight := tui.TerminalDimensions()
	centerX, centerY := termWidth/2, termHeight/2
	rectWidth, rectHeight := 25, 3
	x1, y1 := centerX-rectWidth/2, centerY-rectHeight/2
	x2, y2 := centerX+rectWidth/2, centerY+rectHeight/2
	if rectHeight%2 != 0 {
		y2++
	}
	return x1, y1, x2, y2
}

func newCounterBox() *widgets.Paragraph {
	x1, y1, x2, y2 := rectDimensions(25, 3)
	p := widgets.NewParagraph()
	p.SetRect(x1, y1, x2, y2)
	p.TextStyle.Fg = tui.ColorWhite
	return p
}
