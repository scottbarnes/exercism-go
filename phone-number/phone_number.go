package phonenumber

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// Number cleans a user-entered phone number by returning just the number
// and no additional characters. It also removes any +1 country code.
func Number(phoneNumber string) (string, error) {
	var sb strings.Builder

	// Build a new string with only numerals
	for _, r := range phoneNumber {
		if unicode.IsNumber(r) {
			sb.WriteRune(r)
		}
	}

	// Write out the string and strip and +1 country code present.
	cleanNumber := sb.String()
	cleanNumber = strings.TrimLeft(cleanNumber, "10")

	// Validity checks
	switch {
	// Check for numbers that are too long or too short.
	case len(cleanNumber) != 10:
		return "", errors.New("invalid number of digits")
	// Check for valid exchange prefixes.
	case cleanNumber[3] == '0' || cleanNumber[3] == '1':
		return "", errors.New("exchange code cannot start with 1 or 0")
	}

	return cleanNumber, nil
}

// AreaCode returns a valid area code for a user-entered phone number.
func AreaCode(phoneNumber string) (string, error) {
	// Number cleans, so just check for 10 digits and return the first 3.
	number, err := Number(phoneNumber)
	if len(number) == 10 {
		return number[:3], err
	}

	return "", errors.New("invalid number")
}

// Format returns a user-entered phone number in (123) 456-7890 format.
func Format(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if len(number) == 10 {
		return fmt.Sprintf("(%v) %v-%v", number[:3], number[3:6], number[6:]), err
	}

	return "", errors.New("invalid number")

}
