package cards

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// a the given index existed in the slice or not.
func GetItem(slice []uint8, index int) (uint8, bool) {
	// If len(slice) > index, then index must be within the slice.
	if len(slice) > index && index >= 0 {
		return slice[index], true
	} else {
		return 0, false
	}
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range it is be appended.
func SetItem(slice []uint8, index int, value uint8) []uint8 {
	if len(slice) > index && index >= 0 {
		slice[index] = value
		return slice
	} else {
		r := append(slice, value)
		return r
	}
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
	var s []int
	for i := 0; i < length; i++ {
		s = append(s, value)
	}
	return s

}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if len(slice) > index && index >= 0 {
		return append(slice[:index], slice[index+1:]...)
	} else {
		return slice
	}
	// The above obviates the need for the following:
	// switch {
	// case slice == nil:
	// 	return slice
	// case len(slice) > cap(slice):
	// 	return slice
	// case len(slice) <= index:
	// 	return slice
	// case index < 0:
	// 	return slice
	// default:
	// 	return append(slice[:index], slice[index+1:]...)
	// }
}
