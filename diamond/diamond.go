package diamond

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"strings"
)

// special line for A?
// before, letter, between, letter, after

// getFirstAndLast returns the first row (identical to the last).
func getFirstAndLast(rows int) string {
	var b strings.Builder
	for i := 0; i < rows; i++ {
		b.WriteString(" ")
	}

	padding := b.String()

	if rows > 0 {
		return fmt.Sprintf("%s%s%s", padding, "A", padding)
	} else {
		return "A"
	}
}

// 0 .····A····.
// 1 .···B·B···.
// 2 .··C123C··.
// 3 .·D12345D·.
// 4 .E1234567E.
// 5 F123456789F
// 6 .E1234567E.
// 7 .·D·····D·.
// 8 .··C···C··.
// 9 .···B·B···.
// 0 .····A····.

// getSpaces returns a string of spaces of num length.
func getSpaces(num int) string {
	var b strings.Builder
	for i := 0; i < num; i++ {
		b.WriteString(" ")
	}
	return b.String()
}

// getEdgePadding gets the edge padding based on the 0-indexed row number
func getEdgePadding(row int, rows int) string {
	// The edge padding is always equal total rows - the row number.
	pads := rows - row
	return getSpaces(pads)
}

func getForwardSequence(rows int) []string {
	firstAndLast := getFirstAndLast(rows)
	initialLetter := 'A'
	middleSpaceCount := 1
	var b strings.Builder

	for i := 1; i <= rows; i++ {
		letter := string(byte(initialLetter) + byte(i))
		middleSpaces := getSpaces(middleSpaceCount)
		middleSpaceCount += 2
		middleRow := letter + middleSpaces + letter
		edgePadding := getEdgePadding(i, rows)
		b.WriteString(edgePadding + middleRow + edgePadding + "\n")
	}

	buf := bytes.Buffer{}
	buf.WriteString(b.String())
	scanner := bufio.NewScanner(&buf)

	result := []string{firstAndLast}
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}

func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", errors.New("invalid letter")
	}

	rowsByte := char - 'A'
	rows := int(rowsByte)

	result := ""

	// Get forward sequence
	forwardsequence := getForwardSequence(rows)
	for _, line := range forwardsequence {
		result = result + line + "\n"
	}

	// Get backward sequence
	letterCount := len(forwardsequence)
	for i := letterCount - 2; i >= 0; i-- {
		result = result + forwardsequence[i] + "\n"
	}

	result = strings.TrimRight(result, "\n")

	return result, nil
}
