package main

import (
	"image/color"
	"math/rand"
	"time"
)

const (
	gridNumRows = 10
	tileWidth   = 10
	tileHeight  = 10
)

type GameType struct {
	NumCols    int
	NumRows    int
	NumObjects int
	ObjectFreq map[int]int
}

var game GameType

func (g *GameType) init(rows int) {
	g.NumRows = rows
	g.NumCols = rows*2 - 1
	g.NumObjects = g.NumRows * g.NumCols
	g.ObjectFreq = make(map[int]int, 0)
	for i := 0; i < g.NumObjects; i++ {
		g.ObjectFreq[i] = 0
	}
}

func (g *GameType) startCol() int {
	return (g.NumCols + 1) / 2
}

func (g *GameType) objectNum(row, col int) int {
	return (row * g.NumCols) + col
}

type GridType struct {
	Width, Height    int // pixels wide and high
	NumRows, NumCols int // how many tiles - rows and columns
	grid             [][]color.Color
	triangle         map[int][]int
}

var Grid GridType

func Init() {
	Grid.init(gridNumRows, tileWidth, tileHeight)
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
	gt.Height = numRows * tileHeight

	gt.NumCols = numRows*2 + 1
	gt.Width = gt.NumCols * tileWidth

	gt.triangle = make(map[int][]int)
	for r := 0; r < numRows; r++ {
		for r1 := -1 * r; r1 <= r; r1 += 2 {
			gt.triangle[r] = append(gt.triangle[r], r1)
		}
	}

	gt.grid = make([][]color.Color, gt.Width)
	for x := range gt.grid {
		gt.grid[x] = make([]color.Color, gt.Height)
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
