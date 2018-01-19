package boundary

import (
	"math"
)

//sumOfInts returns the sum of the arguments if they are integers; i.e, have no fractional part.
func sumOfInts(floats ...float64) (sum int) {
	for _, x := range floats {
		if math.Floor(x) == x {
			// this seems good, but math.Floor(±∞) == math.Floor(±∞)
			sum += int(x)
		}
	}
	return sum
}

const (
	ErrCannotConvertInf Error = "cannot convert ±∞ to an integer"
	ErrCannotConvertNaN Error = "cannot convert NaN to an integer"
)

func sumOfIntsFixed(floats ...float64) (sum int, err error) {
	for _, x := range floats {
		switch {
		case math.IsInf(x, 0):
			return 0, ErrCannotConvertInf
		case math.IsNaN(x):
			return 0, ErrCannotConvertNaN
		case math.Floor(x) == x:
			sum += int(x)
		default: //pass
		}
	}
	return sum, nil
}
