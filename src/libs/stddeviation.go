package libs

import "math"

func Stddeviation(numbers []int) float64 {
	return math.Sqrt(float64(Variance(numbers)))
}
