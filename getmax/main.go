package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Finding the Max Value in a Slice")

	scores := []int{4, 5, 7, 777, 334, 5, 6778}
	maxVal := math.MinInt

	for _, v := range scores {
		maxVal = int(math.Max(float64(v), float64(maxVal)))
	}

	fmt.Printf("The max value in the given slice is %d\n", maxVal)
}
