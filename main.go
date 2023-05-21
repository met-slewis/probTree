package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"image/color"
)

func main() {
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
			updateImage()
		}
	}()

	myWindow.Resize(fyne.NewSize(float32(Grid.Width), float32(Grid.Height)))
	myWindow.ShowAndRun()
}
