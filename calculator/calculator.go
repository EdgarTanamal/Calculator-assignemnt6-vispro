package calculator

import (
	"errors"
	"fmt"
	"mathops/mathops"
	"net/http"
	"strconv"
)

func Calculate(a, b float64, operator rune) (float64, error) {
	var result float64
	var err error

	switch operator {
	case '+':
		result = mathops.Addition(a, b)
	case '-':
		result = mathops.Subtraction(a, b)
	case '*':
		result = mathops.Multiplication(a, b)
	case '/':
		if b == 0 {
			return 0, errors.New("division by zero is not allowed")
		}

		result, err = mathops.Division(a, b)
		if err != nil {
			return 0, err
		}
	case '^':
		result, err = mathops.Power(a, b)
		if err != nil {
			return 0, err
		}
	default:
		return 0, errors.New("invalid operator")
	}

	return result, nil
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	number1Str := r.URL.Query().Get("number1")
	number2Str := r.URL.Query().Get("number2")
	operator := r.URL.Query().Get("operator")

	number1, err := strconv.ParseFloat(number1Str, 64)
	if err != nil {
		http.Error(w, "Invalid input for number1", http.StatusBadRequest)
		return
	}

	number2, err := strconv.ParseFloat(number2Str, 64)
	if err != nil {
		http.Error(w, "Invalid input for number2", http.StatusBadRequest)
		return
	}

	result, err := Calculate(number1, number2, []rune(operator)[0])
	if err != nil {
		http.Error(w, fmt.Sprintf("Error in calculation: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	resultStr := strconv.FormatFloat(result, 'f', -1, 64)
	fmt.Fprintf(w, "Hasil: %s", resultStr)
}
