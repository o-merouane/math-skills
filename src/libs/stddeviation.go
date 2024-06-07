package libs

import "math"

func Stddeviation(numbers []int) float64 {

	res := math.Sqrt(float64(Variance(numbers)))

	return res
}
