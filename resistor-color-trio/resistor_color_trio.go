package resistorcolortrio

import (
	"fmt"
	"strconv"
	"strings"
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

func Label(colors []string) string {
	firstColor := colorMap[colors[0]]
	secondColor := colorMap[colors[1]]
	thirdColor := colorMap[colors[2]] - 1

	// zeros := thirdColor - 1

	zerothPlace := ""

	if thirdColor >= 0 {
		for i := 0; i <= thirdColor; i++ {
			zerothPlace += "0"
		}
	}

	fmt.Println(colors, zerothPlace)

	number := fmt.Sprintf("%d%d%s", firstColor, secondColor, zerothPlace)

	suffix := "ohms"
	numInt, _ := strconv.Atoi(number)
	// Handle conversions
	if len(number) >= 9 {
		suffix = "gigaohms"
		numInt = numInt / 1000000000
	} else if len(number) >= 7 {
		suffix = "megaohms"
		numInt = numInt / 1000000
	} else if len(number) >= 4 {
		suffix = "kiloohms"
		numInt = numInt / 1000
	}
	number = strconv.Itoa(numInt)

	// Remove leading 0s.
	if len(number) >= 2 {
		number = strings.TrimPrefix(number, "0")
	}

	return fmt.Sprintf("%s %s", number, suffix)
}
