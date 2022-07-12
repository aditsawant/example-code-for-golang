package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("FizzBuzz Problem")

	n := getInput()

	if n%5 == 0 && n%3 == 0 {
		fmt.Println("fizz buzz")
	} else if n%3 == 0 {
		fmt.Println("fizz")
	} else if n%5 == 0 {
		fmt.Println("buzz")
	} else {
		fmt.Println(n)
	}
}

func getInput() int {
	fmt.Print("Please enter an integer: ")
	reader := bufio.NewReader(os.Stdin)
	nString, _ := reader.ReadString('\n')
	nString = strings.TrimSpace(nString)
	n, err := strconv.ParseInt(nString, 10, 64)
	if err != nil {
		fmt.Println("Please enter a valid integer!")
		panic(err)
	}
	return int(n)
}
