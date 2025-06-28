package mathutils

import (
	"errors"
	"math"
)

// Basic arithmetic operations

// Add returns the sum of two numbers
func Add(a, b float64) float64 {
	return a + b
}

// Subtract returns the difference of two numbers
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply returns the product of two numbers
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide returns the quotient of two numbers
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Advanced math operations

// Power returns base raised to the power of exponent
func Power(base, exponent float64) float64 {
	return math.Pow(base, exponent)
}

// SquareRoot returns the square root of a number
func SquareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("square root of negative number")
	}
	return math.Sqrt(x), nil
}

// Abs returns the absolute value of a number
func Abs(x float64) float64 {
	return math.Abs(x)
}

// Factorial returns the factorial of a non-negative integer
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial of negative number")
	}
	if n == 0 || n == 1 {
		return 1, nil
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result, nil
}