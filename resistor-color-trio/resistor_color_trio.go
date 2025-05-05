package resistorcolortrio

import (
	"fmt"
	"math"
)

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").

var colorMap = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

type Ohm struct {
	Value int
	Unit  string
}

var ohms = []Ohm{
	{1_000_000_000, "gigaohms"},
	{1_000_000, "megaohms"},
	{1_000, "kiloohms"},
}

// Return the combined number. E.g. 3, 3, 0 becomes 33, and 3, 3, 1 becomes 330.
func getCombinedNumber(first, second, third int) int {
	return (first*10 + second) * int(math.Pow10(third))
}

// Format the resistance value with the appropriate unit and value.
func formatResistance(resistence int, ohms []Ohm) string {
	value := resistence
	unit := "ohms"

	for _, ohm := range ohms {
		if resistence >= ohm.Value {
			unit = ohm.Unit
			value = resistence / ohm.Value
			break
		}
	}

	return fmt.Sprintf("%d %s", value, unit)
}

func Label(colors []string) string {
	firstColor := colorMap[colors[0]]
	secondColor := colorMap[colors[1]]
	thirdColor := colorMap[colors[2]]

	number := getCombinedNumber(firstColor, secondColor, thirdColor)

	return formatResistance(number, ohms)
}
