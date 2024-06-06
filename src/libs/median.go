package libs

func Median(num []int) int {

	var x int
	k := 0

	for k < len(num)-2 {
		for i := 0; i < len(num)-1; i++ {
			if num[i] > num[i+1] {
				x = num[i+1]
				num[i+1] = num[i]
				num[i] = x
			}
		}
		k++
	}

	var res float64
	if len(num)%2 != 0 {
		res = float64(num[len(num)/2])
	} else {
		res = float64(num[len(num)/2] + num[len(num)/2-1]) / 2.0
	}

	s := Round(res)
	return s
}
