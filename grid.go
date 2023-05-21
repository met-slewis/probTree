package main

import (
	"image/color"
	"math/rand"
	"time"
)

const (
	numRows    = 5
	tileWidth  = 10
	tileHeight = 10
)

type GridType struct {
	Width, Height    int
	NumRows, NumCols int
	grid             [][]color.Color
	triangle         map[int][]int
}

var Grid GridType

func Init() {
	Grid.init(5, tileWidth, tileHeight)
	rand.Seed(time.Now().UnixNano())
	for x := 0; x < Grid.Width; x++ {
		for y := 0; y < Grid.Height; y++ {
			xCol := uint8(0xff * x / Grid.Width)
			yCol := uint8(0xff * y / Grid.Height)
			col := color.RGBA{R: xCol, G: yCol, A: 0xFF}
			Grid.grid[x][y] = col
		}
	}
}

func (gt *GridType) init(numRows, tileWidth, tileHeight int) {
	gt.NumRows = numRows
	gt.NumCols = numRows*2 + 1
	gt.Width = gt.NumCols * tileWidth
	gt.Height = numRows * tileHeight
	gt.triangle = make(map[int][]int)

	for r := 0; r < numRows; r++ {
		for r1 := -1 * r; r1 <= r; r1 += 2 {
			gt.triangle[r] = append(gt.triangle[r], r1)
		}
	}

}

func (gt *GridType) setTile(row, col int, color color.Color) {
	rowPixel := row * tileHeight
	colPixel := col * tileWidth
	for r := rowPixel; r < rowPixel+tileHeight; r++ {
		for c := colPixel; c < colPixel+tileWidth; c++ {
			gt.grid[c][r] = color
		}
	}
}
