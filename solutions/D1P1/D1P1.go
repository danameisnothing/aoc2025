package d1p1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/danameisnothing/aoc2025/types/typesD1P1"
)

const (
	StartNum     int = 50
	MaxThreshold int = 99
	MinThreshold int = 0
)

func Solve(input string) {
	dial := typesD1P1.NewDial(StartNum, MaxThreshold, MinThreshold)

	var timesIsZero int

	for v := range strings.SplitSeq(input, "\n") {
		if strings.TrimSpace(v) == "" {
			continue
		}

		rotation, err := typesD1P1.Categorize(string(v[0]))
		if err != nil {
			panic(err)
		}

		num, err := strconv.Atoi(v[1:])
		if err != nil {
			panic(err)
		}

		switch rotation {
		case typesD1P1.Left:
			dial.DoLeft(num)
		case typesD1P1.Right:
			dial.DoRight(num)
		}

		if dial.GetNumber() == 0 {
			timesIsZero++
		}
	}

	fmt.Printf("%v\n", timesIsZero)
}
