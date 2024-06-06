package libs

func Average(num []int) int{
	
	var sum int

	for i := 0; i < len(num); i++ {

		sum += (num[i])
	}
	avg := float64(sum) / float64(len(num))

	s := Round(avg)

	return s
}