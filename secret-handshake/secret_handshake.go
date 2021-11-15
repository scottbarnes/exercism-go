package secret

// v2. Thanks @explodes.
// Handshake outputs the order of the secret handshake, using a binary code.
func Handshake(code uint) []string {
	var output = []string{}
	codes := []string{"wink", "double blink", "close your eyes", "jump"}

	for power := 0; power <= 3; power++ {
		if code&(1<<power) > 0 {
			output = append(output, codes[power])
		}

	}
	if code&(1<<4) > 0 {
		for i, j := 0, len(output)-1; i < j; i, j = i+1, j-1 {
			output[i], output[j] = output[j], output[i]
		}
	}

	// if code&(1<<0) > 0 {
	// 	output = append(output, "wink")
	// }
	// if code&(1<<1) > 0 {
	// 	output = append(output, "double blink")
	// }
	// if code&(1<<2) > 0 {
	// 	output = append(output, "close your eyes")
	// }
	// if code&(1<<3) > 0 {
	// 	output = append(output, "jump")
	// }
	// if code&(1<<4) > 0 {
	// 	for i, j := 0, len(output)-1; i < j; i, j = i+1, j-1 {
	// 		output[i], output[j] = output[j], output[i]
	// 	}
	// }

	return output
}
