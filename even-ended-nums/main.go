package main

import (
	"fmt"
)

func main() {
	fmt.Println("Even Ended nums")

	count := 0

	for i := 1000; i <= 9999; i++ {
		for j := i; j <= 9999; j++ {
			prod := i * j
			prodStr := fmt.Sprintf("%d", prod)
			if prodStr[0] == prodStr[len(prodStr)-1] {
				count++
			}
		}
	}

	fmt.Printf("There are %d even ended numbers that are also a product of two four digit nums\n", count)
}
