package main

import (
	"fmt"
	"mathskills/src/libs"
	"os"
)

func main() {

	numbers := libs.ReadFile(os.Args[1])

	var0 := libs.Round(libs.Average(numbers))
	var1 := libs.Round(libs.Variance(numbers))
	var2 := libs.Round(libs.Stddeviation(numbers))

	fmt.Println("Average:", var0)
	fmt.Println("Median:", libs.Median(numbers))
	fmt.Println("Variance:", var1)
	fmt.Println("Standard Deviation:", var2)

}
