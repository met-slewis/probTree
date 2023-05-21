package main

import (
	"fmt"
	"testing"
)

func TestTriangle(t *testing.T) {
	for r := 0; r < 6; r++ {
		for r1 := -1 * r; r1 <= r; r1 += 2 {
			fmt.Printf("%d ", r1)
		}
		fmt.Println("")
	}
}
