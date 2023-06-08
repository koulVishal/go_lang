package math

import "testing"

// Test case for the Add function
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) returned %d, expected %d", result, expected)
	}
}

// Test case for the Subtract function
func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	expected := 2

	if result != expected {
		t.Errorf("Subtract(5, 3) returned %d, expected %d", result, expected)
	}
}

// Test case for the Multiply function
func TestMultiply(t *testing.T) {
	result := Multiply(2, 3)
	expected := 6

	if result != expected {
		t.Errorf("Multiply(2, 3) returned %d, expected %d", result, expected)
	}
}

// Test case for the Divide function
func TestDivide(t *testing.T) {
	result := Divide(10, 2)
	expected := 5

	if result != expected {
		t.Errorf("Divide(10, 2) returned %d, expected %d", result, expected)
	}
}
