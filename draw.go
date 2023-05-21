package main

import (
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
