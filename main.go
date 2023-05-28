package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	game.init(15)

	myApp := app.New()
	myWindow := myApp.NewWindow("Probability Tree")

	content := gridLayoutContainer()

	myWindow.SetContent(content)

	// Start a goroutine to update the drawing continuously
	go func() {
		row := 0
		col := game.startCol()
		for {
			row, col = nextMove(content, row, col)
			content.Refresh()
			time.Sleep(10 * time.Millisecond)
			if row >= game.NumRows {
				row = 0
				col = game.startCol()
			}
		}
	}()

	myWindow.Resize(fyne.NewSize(400, 400))
	myWindow.ShowAndRun()
}
