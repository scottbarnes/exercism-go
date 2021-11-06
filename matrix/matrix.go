package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// [][]int is the same as map[int][]int. Takes an integer representing a row and returns
// slice of ints for that row
type Matrix [][]int

func New(s string) (*Matrix, error) {
	splitString := strings.Split(s, "\n")
	m := Matrix{}

	// Convert the string to []ints and then append to the Matrix.
	for _, v := range splitString {
		slice, err := stringToIntSlice(v)
		if err != nil {
			return &m, err
		}
		m = append(m, slice)
	}

	if ok, err := m.isMatrixValid(); !ok {
		return &m, err
	}
	return &m, nil
}

func (m *Matrix) isMatrixValid() (bool, error) {
	// All slices must be the same length. Match all slice lengths against the first one.
	targetLength := len((*m)[0])
	for i := range *m {
		if len((*m)[i]) != targetLength {
			return false, errors.New("sets of different lengths")
		}
	}
	return true, nil
}

// stringToIntSlice takes a string and returns an []int or an error.
func stringToIntSlice(s string) ([]int, error) {
	var result = []int{}

	// Split the string into []string. Then for each string in []string, convert it to an integer,
	// then append it to our []int for returning.
	fields := strings.Fields(s)
	for _, numString := range fields {
		numInt, err := strconv.Atoi(numString)
		if err != nil {
			return []int{}, err
		}
		result = append(result, numInt)
	}
	return result, nil
}

func (m *Matrix) Cols() [][]int {
	// Get the length of the elements so we know how many new subslices there should be,
	// as for each element there will be a new subslice.
	sliceLen := len((*m)[0])
	col := make([][]int, sliceLen)

	// For each new slice, make an inner slice with the length of the matrix.
	// Because x and y are being transposed, so to speak.
	for i := 0; i < len((*m)[0]); i++ {
		col[i] = make([]int, len(*m))
	}

	// Take from the columns and add to the rows. The 0th elements from each row
	// constitute the 0th, 1st, 2nd, ... nth of each row.
	for x, rows := range *m {
		for i, element := range rows {
			col[i][x] = element
		}
	}
	return col
}

func (m *Matrix) Rows() [][]int {
	// Create as many rows as the length of the matrix, then fill in each with a subslice
	row := make([][]int, len(*m))
	for i, mRow := range *m {
		row[i] = make([]int, len(mRow))
	}
	for i, v := range *m {
		copy(row[i], v)
	}

	return row
}

func (m *Matrix) Set(row, column, value int) bool {
	r := m.Rows()
	if row < 0 || column < 0 {
		return false
	}
	maxLength := len(r[0])
	if row >= maxLength || column >= maxLength {
		return false
	}
	(*m)[row][column] = value // what is the parenthesis around *m doing?
	return true
}
