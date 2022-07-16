// Following this blog: https://blog.logrocket.com/concurrency-patterns-golang-waitgroups-goroutines/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("Go Routines")

	wg := new(sync.WaitGroup)
	// fmt.Printf("%v: %T", wg, wg)

	wg.Add(2)
	go randInt(wg, "First: ", 10, 2)
	go randInt(wg, "Second: ", 3, 2)
	wg.Wait()
}

func randInt(wq *sync.WaitGroup, fn string, lim int, delay int) {
	defer wq.Done()
	for i := 1; i <= lim; i++ {
		fmt.Println(fn, rand.Intn(i))
		time.Sleep(time.Duration(int(time.Second) * delay))
	}
}
