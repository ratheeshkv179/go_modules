package numbers

import (
	"errors"
)

// Sum of input values
func Sum(a int, b int) int {
	return a + b
}

// Difference of input values
func Difference(a int, b int) (int, error) {
	if a >= b {
		return a - b, nil
	}
	return 0, errors.New("Invalid Input")

}

func Product(a int, b int) int {
	return a * b
}
