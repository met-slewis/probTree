package main

import (
	"image/color"
	"math/rand"
	"time"
)

func updateImage() {
	for row := 0; row < Grid.NumRows; row++ {
		for col := 0; col < Grid.NumCols; col++ {
			r := uint8(rand.Intn(0xFF))
			g := uint8(rand.Intn(0xFF))
			b := uint8(rand.Intn(0xFF))
			colr := color.RGBA{R: r, G: g, B: b, A: 0xFF}
			Grid.setTile(row, col, colr)
		}
	}
	time.Sleep(500 * time.Millisecond)
}
