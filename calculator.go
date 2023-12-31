// Package calculator does simple calculations
package calculator

import (
	"errors"
	"math"
)

func Add(a, b float64) float64 {
	return a + b
}

func Substract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero not allawed")
	}
	return a / b, nil
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("square root of negative number is not possible")
	}
	return math.Sqrt(a), nil
}
