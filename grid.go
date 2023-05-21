package main

import (
	"image/color"
	"math/rand"
	"time"
)

const (
	height     = 400
	width      = 400
	tileWidth  = 10
	tileHeight = 10
)

type GridType struct {
	Width, Height    int
	NumRows, NumCols int
	grid             [width][height]color.Color
}

var Grid GridType

func Init() {
	Grid.init(width, height)
	rand.Seed(time.Now().UnixNano())
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			xCol := uint8(0xff * x / width)
			yCol := uint8(0xff * y / height)
			col := color.RGBA{R: xCol, G: yCol, A: 0xFF}
			Grid.grid[x][y] = col
		}
	}
}

func (gt *GridType) init(width, height int) {
	gt.Width = width
	gt.Height = height
	gt.NumRows = height / tileHeight
	gt.NumCols = width / tileWidth
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
