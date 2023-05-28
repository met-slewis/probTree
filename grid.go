package main

import "image/color"

type GameType struct {
	NumCols       int
	NumRows       int
	NumObjects    int
	ObjectFreq    map[int]int
	ColorGradient []color.Color
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
	g.computeGradient(128)
}

func (g *GameType) startCol() int {
	return ((g.NumCols - 1) / 2) + 1
}

func (g *GameType) objectNum(row, col int) int {
	return (row * g.NumCols) + col - 1
}

func (gme *GameType) computeGradient(numSteps int) {
	blue := [3]int{0, 0, 255}
	red := [3]int{255, 0, 0}

	gme.ColorGradient = make([]color.Color, 0)

	for step := 0; step < numSteps; step++ {
		r := blue[0] + ((red[0]-blue[0])/(numSteps-1))*step
		g := blue[1] + ((red[1]-blue[1])/(numSteps-1))*step
		b := blue[2] + ((red[2]-blue[2])/(numSteps-1))*step

		gme.ColorGradient = append(gme.ColorGradient, color.RGBA{
			R: uint8(r),
			G: uint8(g),
			B: uint8(b),
			A: 0xFF,
		})
	}
}
