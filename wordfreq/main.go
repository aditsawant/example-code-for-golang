package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Printing the word freq map")

	sentence := `That's why I need a one dance
	Got a Hennessy in my hand
	One more time 'fore I go
	Higher powers taking a hold on me
	I need a one dance
	Got a Hennessy in my hand
	One more time 'fore I go
	Higher powers taking a hold on me`

	// fmt.Println(sentence)
	freq := make(map[string]int)
	// words := strings.Split(sentence, " ")
	words := strings.Fields(sentence)

	// fmt.Println(len(words))
	for _, v := range words {
		freq[strings.ToLower(v)]++
	}

	for k, v := range freq {
		fmt.Printf("%v: %v\n", k, v)
	}
}
