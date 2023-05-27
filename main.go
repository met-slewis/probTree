package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"image/color"
	"math/rand"
	"time"
)

func mainYY() {
	gridLayout()
}

func mainXX() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Real-Time Drawing")

	Init()
	canvas := canvas.NewRasterWithPixels(func(x, y, width, height int) color.Color {
		return Grid.grid[x][y]
	})

	content := container.New(layout.NewMaxLayout(), canvas)
	myWindow.SetContent(content)

	// Start a goroutine to update the drawing continuously
	go func() {
		for {
			canvas.Refresh() // Redraw the canvas
			// Perform any necessary updates to the drawing here
			updateImageTmp()
		}
	}()

	myWindow.Resize(fyne.NewSize(float32(Grid.Width), float32(Grid.Height)))
	myWindow.ShowAndRun()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	game.init(4)

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
			time.Sleep(250 * time.Millisecond)
			if row > game.NumRows {
				row = 0
				col = game.startCol()
			}
		}
	}()

	myWindow.Resize(fyne.NewSize(400, 400))
	myWindow.ShowAndRun()
}
