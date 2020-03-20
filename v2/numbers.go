package numbers

import (
	"errors"
)

// Sum of input values
func Sum(a ...) int {

	sum := 0
	for _, v := range a{
		sum += v
	}
	return sum
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
