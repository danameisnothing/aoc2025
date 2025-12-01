package typesD1P2

import (
	"fmt"
	"math"
)

type Rotation uint8

const (
	Left = iota
	Right
)

var rotationMapping = map[string]Rotation{
	"L": Left,
	"R": Right,
}

func Categorize(direction string) (Rotation, error) {
	// https://stackoverflow.com/a/2050629
	val, ok := rotationMapping[direction]
	if !ok {
		return math.MaxUint8, fmt.Errorf("no match for direction char %s", direction)
	}

	return val, nil
}
