// Twelve outputs 'The Twelve Days of Christmas', by verse or by the entire song.
package twelve

import "fmt"

var DAYS = []string{
	"zeroth",
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

var THINGS = []string{
	"a Partridge in a Pear Tree",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

// Song returns all twelve verses of The Twelve Days of Christmas, with appropriate new lines.
func Song() string {
	var r string
	// Add the verses together and add a new line at the end of each one, save for the last verse.
	for i := 1; i <= 12; i++ {
		r += Verse(i)
		if i < 12 {
			r += "\n"
		}
	}
	return r
}

// Verse takes a verse integer and returns the correct verse.
func Verse(n int) string {
	v := fmt.Sprintf("On the %v day of Christmas my true love gave to me: ", DAYS[n])
	for i := n - 1; i >= 0; i-- {
		v += THINGS[i] + suffixer(i)
	}
	return v
}

// suffixer takes a verse integer and returns the appropriate punctuation/suffix to follow that verse.
func suffixer(i int) string {
	switch i {
	case 0:
		return "."
	case 1:
		return ", and "
	default:
		return ", "
	}
}
