package d1p1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/danameisnothing/aoc2025/types/typesD1"
)

const (
	StartNum     int = 50
	MaxThreshold int = 99
	MinThreshold int = 0
)

func Solve(input string) {
	dial := typesD1.NewDial(StartNum, MaxThreshold, MinThreshold)

	var timesIsZero int

	for _, v := range strings.Split(input, "\n") {
		if strings.TrimSpace(v) == "" {
			continue
		}

		rotation, err := typesD1.Categorize(string(v[0]))
		if err != nil {
			panic(err)
		}

		num, err := strconv.Atoi(v[1:])
		if err != nil {
			panic(err)
		}

		switch rotation {
		case typesD1.Left:
			dial.DoLeft(num)
		case typesD1.Right:
			dial.DoRight(num)
		}

		if dial.GetNumber() == 0 {
			timesIsZero++
		}
	}

	fmt.Printf("%v\n", timesIsZero)
}
