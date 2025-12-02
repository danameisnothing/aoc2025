package d2p2

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

			var lastMatch string
			for i := range len(numStr) {
				// assuming every element is the same
				matches := regexp.MustCompile(numStr[i+1:]).FindAll([]byte(numStr), -1)
				if len(string(matches[0]))*len(matches) == len(numStr) {
					lastMatch = numStr[i+1:]
					break
				}
			}

			// HACKHACK:
			if lastMatch == "" {
				continue
			}

			var sb strings.Builder
			for {
				if sb.Len() >= len(numStr) {
					break
				}

				sb.WriteString(lastMatch)
				if sb.String() == numStr {
					fmt.Printf("MATCH: %s\n", sb.String())
					toAdd, err := strconv.Atoi(sb.String())
					if err != nil {
						panic(err)
					}

					total += toAdd
				}
			}
		}
	}

	fmt.Printf("%d\n", total)
}
