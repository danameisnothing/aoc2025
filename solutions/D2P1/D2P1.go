package d2p1

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	RegexMatchRanges string = `\d+`
)

func Solve(input string) {
	var total int

	regexpMatchDigitDash := regexp.MustCompile(RegexMatchRanges)
	for v := range strings.SplitSeq(input, ",") {
		rangeNums := regexpMatchDigitDash.FindAll([]byte(v), 2)

		lower, err := strconv.Atoi(string(rangeNums[0]))
		if err != nil {
			panic(err)
		}

		upper, err := strconv.Atoi(string(rangeNums[1]))
		if err != nil {
			panic(err)
		}

		for i := lower; i <= upper; i++ {
			numStr := strconv.Itoa(i)
			// check if the length is even
			if len(numStr)%2 != 0 {
				continue
			}

			var lastMatch string
			for i := range len(numStr) / 2 {
				//fmt.Printf("-> %s, %s\n", numStr[i+1:], numStr[:i+1])
				if !strings.Contains(numStr, numStr[i+1:]) {
					break
				}

				lastMatch = numStr[i+1:]
			}

			comb := fmt.Sprintf("%s%s", lastMatch, lastMatch)
			if comb == numStr {
				toAdd, err := strconv.Atoi(comb)
				if err != nil {
					panic(err)
				}

				total += toAdd
			}
		}
	}

	fmt.Printf("%d\n", total)
}
