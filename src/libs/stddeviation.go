package libs

import "math"

func Stddeviation(numbers []int) int {

	res := math.Sqrt(float64(Variance(numbers)))

	s := Round(res)

	return s
}
