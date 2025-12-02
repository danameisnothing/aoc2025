package d1p2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/danameisnothing/aoc2025/types/typesD1P2"
)

const (
	StartNum     int = 50
	MaxThreshold int = 99
	MinThreshold int = 0
)

func Solve(input string) {
	dial := typesD1P2.NewDial(StartNum, MaxThreshold, MinThreshold)

	var times int

	for v := range strings.SplitSeq(input, "\n") {
		if strings.TrimSpace(v) == "" {
			continue
		}

		rotation, err := typesD1P2.Categorize(string(v[0]))
		if err != nil {
			panic(err)
		}

		num, err := strconv.Atoi(v[1:])
		if err != nil {
			panic(err)
		}

		switch rotation {
		case typesD1P2.Left:
			times += dial.DoLeft(num)
		case typesD1P2.Right:
			times += dial.DoRight(num)
		}
	}

	fmt.Printf("%v\n", times)
}
