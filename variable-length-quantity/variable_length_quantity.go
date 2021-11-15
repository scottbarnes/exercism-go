package variablelengthquantity

import (
	"errors"
	"strconv"
)

// Decimal, hex, octal, and binary conversion table: https://ascii.cl/conversion.htm
/*
I can treat input as decimal, and compare it to decimal. The following work:
1. 0x7E + 5 = 131 (126 + 5 = 131)
2. 0x40 == 64: true
3. If the input number is <= 127, I guess just add it to the output.
*/

func padLeft(s string, length int, pad string) string {
	for {
		if len(s) < length {
			s = pad + s
		} else {
			return s
		}
	}
}

func EncodeVarint(input []uint32) []byte {
	var output []byte

	// No need to convert, because 127 is the largest number 7 bits can hold.
	for _, n := range input {
		if n <= 127 {
			output = append(output, byte(n))
		} else {
			bs := strconv.FormatUint(uint64(n), 2)
			// break into groups of 7.
			sevenBitBinaries := []byte{}
			var workingString string
			count := 1
			prefix := "0"

			// Count backwards to extract groups of 7 bits.
			for i := len(bs) - 1; i >= 0; i-- {
				workingString = bs[i:i+1] + workingString
				if count%7 == 0 {
					workingString = prefix + workingString
					num, _ := strconv.ParseUint(workingString, 2, 0)
					sevenBitBinaries = append(sevenBitBinaries, byte(num))
					// Only the first byte gets a 0 prefix to terminate the sequence.
					// Also, clear out the string for the next byte, if present.
					prefix = "1"
					workingString = ""
				}
				count++
			}

			// Deal with bytes of fewer than 7 bits by padding zeros on the left
			// as needed, then prefix with 1.
			if len(bs)%7 > 0 {
				workingString = padLeft(workingString, 8, "0")
				workingString = "1" + workingString[1:]
				num, _ := strconv.ParseUint(workingString, 2, 0)
				sevenBitBinaries = append(sevenBitBinaries, byte(num))
			}

			for i, j := 0, len(sevenBitBinaries)-1; i < j; i, j = i+1, j-1 {
				sevenBitBinaries[i], sevenBitBinaries[j] = sevenBitBinaries[j], sevenBitBinaries[i]
			}
			// return sevenBitBinaries
			output = append(output, sevenBitBinaries...)
		}
	}

	return output
}

func DecodeVarint(input []byte) ([]uint32, error) {
	var output []uint32

	// for each byte,  remove the prefixes, and keep track of where the stop bit is.
	// Pull out 7 bit binaries
	if len(input) == 1 {
		bs := strconv.FormatUint(uint64(input[0]), 2)
		if bs[0:1] == "1" && len(bs) == 8 {
			return []uint32{}, errors.New("invalid input")
		}
	}
	// error checking. consolidate with below so not parsing multiple times.

	var sevenBitBinaries string
	for _, n := range input {
		// No need to convert these.
		if n <= byte(127) && n != 0 {
			bs := strconv.FormatUint(uint64(n), 2)
			sevenBitBinaries += bs
			i, err := strconv.ParseUint(sevenBitBinaries, 2, 0)
			if err != nil {
				return []uint32{}, err
			}
			output = append(output, uint32(i))
			sevenBitBinaries = ""
		} else {
			bs := strconv.FormatUint(uint64(n), 2)
			if bs[0:1] == "1" {
				sevenBitBinaries += bs[1:]
			} else {
				if len(bs[1:]) < 7 {
					sevenBitBinaries += bs[1:]
					bs = padLeft(bs, 8, "0")
				}
				sevenBitBinaries += bs[1:]
				i, err := strconv.ParseUint(sevenBitBinaries, 2, 0)
				if err != nil {
					return []uint32{}, err
				}
				output = append(output, uint32(i))
				sevenBitBinaries = ""
			}
		}
	}
	return output, nil
}
