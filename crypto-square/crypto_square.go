package cryptosquare

import (
	"math"
)

// normalizer returns a []rune as only numbers and lower case letters (downcased if necessary)
func normalizer(s string) []rune {
	var normalized []rune
	for _, r := range s {
		switch {
		case 'a' <= r && r <= 'z':
			normalized = append(normalized, r)
		case 'A' <= r && r <= 'Z':
			normalized = append(normalized, r+('a'-'A'))
		case '0' <= r && r <= '9':
			normalized = append(normalized, r)
		}
	}
	return normalized
}

// isNumInt returns true if a number is an integer
func isNumInt(n float64) bool {
	return n == float64(int(n))
}

// findGrid takes a integer string length and returns [2]int of columns and rows, where
// r * c >= length(string),
// and c >= r,
// and c - r <= 1.
// Happily, we can use the square root of length(string), and some rounding, to find this.
// Taken together, sqrt(9) = 3 = 3x3 grid. Sqrt(12) = 3.46 (round down) = 3x4 grid, and
// sqroot(33) = 5.74 (round up) = 6x6 grid.
func findGrid(n int) (c, r int) {
	rows := math.Sqrt(float64(n))
	// If math.Sqrt returns an int, grid = n * n. If not, grid = n+1 * n.
	if isNumInt(rows) {
		columns := rows
		return int(columns), int(rows)
	}

	// For non-ints, determine whether to add 1 to just columns, or both columns and rows.
	// If the square root rounds down, just add one to the row. But if it rounds up,
	// increase both the row and column by one.
	if math.Round(rows) < rows {
		columns := rows + 1
		return int(columns), int(rows)
	} else {
		rows += 1
		columns := rows // because rows is +1, this is now +1 too.
		return int(columns), int(rows)
	}
}

func Encode(pt string) string {
	var encoded []rune
	normalized := normalizer(pt) // get a string of just numbers and lowercased letters
	normalizedLength := len(normalized)
	columns, rows := findGrid(normalizedLength) // determine grid size using square root
	encodedLength := columns * rows
	paddedChunks := encodedLength - normalizedLength // determine necessary padding for grid

	// (If needed), add the padding to the normalized's suffix so that we can read straight
	// through later by grabbing every Nth character, which will include the padded ones.
	if paddedChunks > 0 {
		for i := 0; i < paddedChunks; i++ {
			normalized = append(normalized, ' ')
		}
	}

	// Read every Columnth character Rowth times to encode the string.
	// That is to say, for each column, get every Nth character for the rows, starting at col
	for col := 0; col < columns; col++ {
		offset := col // start counting every Nth character from the column number (e.g. 0, 1, ...n)
		for row := 0; row < rows; row++ {
			encoded = append(encoded, normalized[offset])
			offset += columns // From the column's number, increment by column number
		}
		// Don't add a space to the last row.
		if col < columns-1 {
			encoded = append(encoded, ' ')
		}
	}

	return string(encoded)
}
