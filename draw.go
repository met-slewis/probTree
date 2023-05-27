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

func updateImage() {
	for row := 0; row < Grid.NumRows; row++ {
		for col := 0; col < Grid.NumCols; col++ {
			if col == Grid.NumCols/2 {
				r := uint8(rand.Intn(0xFF))
				g := uint8(rand.Intn(0xFF))
				b := uint8(rand.Intn(0xFF))
				colr := color.RGBA{R: r, G: g, B: b, A: 0xFF}
				Grid.setTile(row, col, colr)
			} else if col == Grid.NumCols/2+1 {
			} else {
				Grid.setTile(row, col, color.Black)
				Grid.setTile(row, col, color.Black)
			}
		}
	}
	time.Sleep(1000 * time.Millisecond)
}

func updateImageTmp() {
	for row := 0; row < Grid.NumRows; row++ {
		for col := 0; col < Grid.NumCols; col++ {
			if row%2 == 0 && col%2 == 0 {
				r := uint8(rand.Intn(0xFF))
				g := uint8(rand.Intn(0xFF))
				b := uint8(rand.Intn(0xFF))
				colr := color.RGBA{R: r, G: g, B: b, A: 0xFF}
				Grid.setTile(row, col, colr)
				//} else if col == Grid.NumCols/2+1 {
				//	Grid.setTile(row, col, color.Black)
			} else {
				Grid.setTile(row, col, color.Black)
			}
		}
	}
	time.Sleep(1000 * time.Millisecond)
}

func drawCircle(row, col int, color color.Color) {

}
func drawBigCircle(row, col int, color color.Color) {

}

func gridLayoutContainer() *fyne.Container {

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
	return grid
}

func nextMove(container *fyne.Container, row, col int) (newRow, newCol int) {
	objNum := game.objectNum(row, col)
	game.ObjectFreq[objNum]++
	container.Objects[objNum] = canvas.NewCircle(color.White)
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
