package libs

func Average(num []int) float64 {

	var sum int

	for i := 0; i < len(num); i++ {

		sum += (num[i])
	}
	avg := float64(sum) / float64(len(num))

	return avg
}
