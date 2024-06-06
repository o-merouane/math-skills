package libs

func Variance(num []int) int {

	var sum int

	for i := 0; i < len(num); i++ {
		x := num[i] - Average(num)
		x *= x
		sum += x

	}

	res := float64(sum) / float64((len(num) - 1))

	s := Round(res)

	return s
}
