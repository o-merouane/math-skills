package libs

import "math"

func Round(num float64) int {

	intPart, fracPart := math.Modf(num)

	if fracPart >= 0.5 {
		return int(intPart + 1)
	}

	return int(intPart)
}
