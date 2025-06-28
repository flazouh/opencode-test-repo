package main

import (
	"fmt"
	"go-math-utils/mathutils"
	"log"
)

func main() {
	fmt.Println("Go Math Utils Demo")
	fmt.Println("==================")

	// Basic arithmetic operations
	fmt.Println("\nBasic Arithmetic Operations:")
	fmt.Printf("Add(10, 5) = %.2f\n", mathutils.Add(10, 5))
	fmt.Printf("Subtract(10, 5) = %.2f\n", mathutils.Subtract(10, 5))
	fmt.Printf("Multiply(10, 5) = %.2f\n", mathutils.Multiply(10, 5))

	result, err := mathutils.Divide(10, 5)
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		fmt.Printf("Divide(10, 5) = %.2f\n", result)
	}

	// Division by zero test
	_, err = mathutils.Divide(10, 0)
	if err != nil {
		fmt.Printf("Divide(10, 0) = Error: %v\n", err)
	}

	// Advanced math operations
	fmt.Println("\nAdvanced Math Operations:")
	fmt.Printf("Power(2, 3) = %.2f\n", mathutils.Power(2, 3))
	fmt.Printf("Abs(-15) = %.2f\n", mathutils.Abs(-15))

	sqrt, err := mathutils.SquareRoot(16)
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		fmt.Printf("SquareRoot(16) = %.2f\n", sqrt)
	}

	// Square root of negative number test
	_, err = mathutils.SquareRoot(-4)
	if err != nil {
		fmt.Printf("SquareRoot(-4) = Error: %v\n", err)
	}

	// Factorial operations
	fmt.Println("\nFactorial Operations:")
	factorial, err := mathutils.Factorial(5)
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		fmt.Printf("Factorial(5) = %d\n", factorial)
	}

	factorial, err = mathutils.Factorial(0)
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		fmt.Printf("Factorial(0) = %d\n", factorial)
	}

	// Factorial of negative number test
	_, err = mathutils.Factorial(-3)
	if err != nil {
		fmt.Printf("Factorial(-3) = Error: %v\n", err)
	}
}