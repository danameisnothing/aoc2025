package main

import (
	"os"

	d1p1 "github.com/danameisnothing/aoc2025/solutions/D1P1"
)

const (
	D1P1Input string = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
)

func main() {
	inp, err := os.ReadFile("inputs/D1/P1/input")
	if err != nil {
		panic(err)
	}

	d1p1.Solve(string(inp))
}
