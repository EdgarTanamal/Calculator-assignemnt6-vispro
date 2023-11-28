package mathops

import (
	"errors"
	"math"
)

func Power(base, exponent float64) (float64, error) {
	result := math.Pow(base, exponent)
	if math.IsNaN(result) || math.IsInf(result, 0) {
		return 0, errors.New("invalid operation: power calculation")
	}
	return result, nil
}
