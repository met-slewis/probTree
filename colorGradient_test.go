package main

import (
	"fmt"
	"testing"
)

func TestComputeGradient(t *testing.T) {
	numSteps := 64
	tgame := GameType{}
	tgame.computeGradient(numSteps)
	fmt.Println(tgame.ColorGradient)
}
