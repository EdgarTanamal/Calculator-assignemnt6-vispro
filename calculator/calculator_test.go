package calculator

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	result, err := Calculate(3, 4, '+')
	if err != nil {
		t.Errorf("Expected no error but got: %v", err)
	}
	expected := 7.0
	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}

	_, err = Calculate(5, 0, '/')
	if err == nil {
		t.Errorf("Expected error but got none")
	}

	result, err = Calculate(2, 3, '^')
	if err != nil {
		t.Errorf("Expected no error but got: %v", err)
	}
	expectedPower := 8.0
	if result != expectedPower {
		t.Errorf("Expected %f but got %f", expectedPower, result)
	}
}
