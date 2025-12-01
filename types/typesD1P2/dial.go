package typesD1P2

import (
	"math"
)

type Dial struct {
	number           int
	maximumThreshold int
	minimumThreshold int
}

func NewDial(startNum, max, min int) *Dial {
	return &Dial{
		number:           startNum,
		maximumThreshold: max,
		minimumThreshold: min,
	}
}

func (d *Dial) GetNumber() int {
	return d.number
}

func (d *Dial) DoLeft(amount int) int {
	var amt int
	// maybe could be done with bit shifts, but i don't know how those work.
	numLeft := amount
	for {
		numLeft = d.number - numLeft

		if numLeft < 0 {
			d.number = d.maximumThreshold + 1
			numLeft = int(math.Abs(float64(numLeft)))
			amt++
			continue
		}

		d.number = numLeft
		return amt
	}
}

func (d *Dial) DoRight(amount int) int {
	var amt int
	// maybe could be done with bit shifts, but i don't know how those work.
	numAdd := amount
	for {
		numAdd = d.number + numAdd

		if numAdd > d.maximumThreshold {
			d.number = 0
			numAdd -= d.maximumThreshold + 1
			amt++
			continue
		}

		d.number = numAdd
		return amt
	}
}
