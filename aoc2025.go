package main

import (
	"os"

	d2p1 "github.com/danameisnothing/aoc2025/solutions/D2P1"
)

func main() {
	inp, err := os.ReadFile("inputs/D2/input")
	if err != nil {
		panic(err)
	}

	d2p1.Solve(string(inp))
}
