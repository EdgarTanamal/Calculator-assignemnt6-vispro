package mathops

import (
	"testing"
)

func TestAddition(t *testing.T) {
	result := Addition(3, 5)
	if result != 8 {
		t.Errorf("Expected 8, but got %f", result)
	}
}

func TestSubtraction(t *testing.T) {
	result := Subtraction(10, 4)
	if result != 6 {
		t.Errorf("Expected 6, but got %f", result)
	}
}

func TestMultiplication(t *testing.T) {
	result := Multiplication(7, 3)
	if result != 21 {
		t.Errorf("Expected 21, but got %f", result)
	}
}

func TestDivision(t *testing.T) {
	result, err := Division(15, 3)
	if err != nil {
		t.Errorf("Division returned an error: %s", err.Error())
	}
	if result != 5 {
		t.Errorf("Expected quotient 5, but got %f", result)
	}
	_, err = Division(10, 0)
	if err == nil {
		t.Errorf("Expected an error for division by zero, but no error returned")
	}
}
