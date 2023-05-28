package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"image/color"
	"math/rand"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func gridLayoutContainer() *fyne.Container {

	numCols := game.NumCols
	numRows := game.NumRows

	tiles := make([]fyne.CanvasObject, 0)
	gray := color.RGBA{R: 0x11, G: 0x11, B: 0x11, A: 0xFF}
	for row := 0; row < numRows; row++ {
		for col := 1; col <= numCols; col++ {
			if (col < (numCols+1)/2-row) || (col > (numCols+1)/2+row) {
				tiles = append(tiles, canvas.NewCircle(gray))
			} else {
				tiles = append(tiles, canvas.NewCircle(color.Black))
			}
		}
	}
	grid := container.New(layout.NewGridLayout(numCols), tiles...)
	return grid
}

func nextMove(container *fyne.Container, row, col int) (newRow, newCol int) {
	objNum := game.objectNum(row, col)
	game.ObjectFreq[objNum]++
	f := uint8(min(game.ObjectFreq[objNum]*2, 127))
	colr := game.ColorGradient[f]
	container.Objects[objNum] = canvas.NewCircle(colr)
	row++
	d := rand.Intn(2)
	if d == 0 {
		col--
	} else {
		col++
	}
	return row, col
}

func gridLayout() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Layout")

	numCols := 9 // must be an odd number
	numRows := (numCols + 1) / 2

	tiles := make([]fyne.CanvasObject, 0)
	gray := color.RGBA{R: 0x11, G: 0x11, B: 0x11, A: 0xFF}
	for row := 0; row < numRows; row++ {
		for col := 1; col <= numCols; col++ {
			if (col < (numCols+1)/2-row) || (col > (numCols+1)/2+row) {
				tiles = append(tiles, canvas.NewCircle(gray))
			} else {
				tiles = append(tiles, canvas.NewCircle(color.Black))
			}
		}
	}
	grid := container.New(layout.NewGridLayout(numCols), tiles...)
	myWindow.SetContent(grid)
	myWindow.Resize(fyne.NewSize(400, 400))
	myWindow.ShowAndRun()
}

func gridLayoutSaved() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Layout")

	numCols := 9 // must be an odd number
	numRows := (numCols + 1) / 2

	tiles := make([]fyne.CanvasObject, 0)
	gray := color.RGBA{R: 0x11, G: 0x11, B: 0x11, A: 0xFF}
	for row := 0; row < numRows; row++ {
		for col := 1; col <= numCols; col++ {
			if (col < (numCols+1)/2-row) || (col > (numCols+1)/2+row) {
				tiles = append(tiles, canvas.NewCircle(gray))
			} else {
				tiles = append(tiles, canvas.NewCircle(color.Black))
			}
		}
	}
	grid := container.New(layout.NewGridLayout(numCols), tiles...)
	myWindow.SetContent(grid)
	myWindow.Resize(fyne.NewSize(400, 400))
	myWindow.ShowAndRun()
}
