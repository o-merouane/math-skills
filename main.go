package main

import (
	"fmt"
	"mathskills/src/libs"
)

func main() {

	numbers := libs.ReadFile("data.txt")

	fmt.Println("Average:", libs.Average(numbers))
	fmt.Println("Median:", libs.Median(numbers))
	fmt.Println("Variance:", libs.Variance(numbers))
	fmt.Println("Standard Deviation:", libs.Stddeviation(numbers))

}
