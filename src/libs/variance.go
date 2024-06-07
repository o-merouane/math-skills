package libs

func Variance(num []int) float64 {

	// Calculate the variance
	var varianceSum float64
	for _, value := range num {
		varianceSum += (float64(value) - Average(num)) * (float64(value) - Average(num))
	}
	variance := varianceSum / float64(len(num))

	return variance

}
