package mathutils

import (
	"testing"
)

// Test basic arithmetic operations
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5.0
	if result != expected {
		t.Errorf("Add(2, 3) = %f; want %f", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	expected := 2.0
	if result != expected {
		t.Errorf("Subtract(5, 3) = %f; want %f", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(4, 3)
	expected := 12.0
	if result != expected {
		t.Errorf("Multiply(4, 3) = %f; want %f", result, expected)
	}
}

func TestDivide(t *testing.T) {
	result, err := Divide(10, 2)
	expected := 5.0
	if err != nil {
		t.Errorf("Divide(10, 2) returned error: %v", err)
	}
	if result != expected {
		t.Errorf("Divide(10, 2) = %f; want %f", result, expected)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := Divide(10, 0)
	if err == nil {
		t.Error("Divide(10, 0) should return error for division by zero")
	}
}

// Test advanced math operations
func TestPower(t *testing.T) {
	result := Power(2, 3)
	expected := 8.0
	if result != expected {
		t.Errorf("Power(2, 3) = %f; want %f", result, expected)
	}
}

func TestSquareRoot(t *testing.T) {
	result, err := SquareRoot(16)
	expected := 4.0
	if err != nil {
		t.Errorf("SquareRoot(16) returned error: %v", err)
	}
	if result != expected {
		t.Errorf("SquareRoot(16) = %f; want %f", result, expected)
	}
}

func TestSquareRootNegative(t *testing.T) {
	_, err := SquareRoot(-4)
	if err == nil {
		t.Error("SquareRoot(-4) should return error for negative input")
	}
}

func TestAbs(t *testing.T) {
	result := Abs(-5)
	expected := 5.0
	if result != expected {
		t.Errorf("Abs(-5) = %f; want %f", result, expected)
	}

	result = Abs(5)
	expected = 5.0
	if result != expected {
		t.Errorf("Abs(5) = %f; want %f", result, expected)
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		input    int
		expected int
		hasError bool
	}{
		{0, 1, false},
		{1, 1, false},
		{5, 120, false},
		{-1, 0, true},
	}

	for _, test := range tests {
		result, err := Factorial(test.input)
		if test.hasError {
			if err == nil {
				t.Errorf("Factorial(%d) should return error", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("Factorial(%d) returned error: %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("Factorial(%d) = %d; want %d", test.input, result, test.expected)
			}
		}
	}
}