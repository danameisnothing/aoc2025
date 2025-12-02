package main

import (
	"os"

	d2p2 "github.com/danameisnothing/aoc2025/solutions/D2P2"
)

func main() {
	inp, err := os.ReadFile("inputs/D2/input")
	if err != nil {
		panic(err)
	}

	d2p2.Solve(string(inp))
}
