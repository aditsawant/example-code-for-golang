// Following this blog for channels in Go: https://blog.logrocket.com/concurrency-patterns-golang-waitgroups-goroutines/

package main

import (
	"fmt"
	"sync"
)

func ChannelMain() {
	fmt.Println("Printing using channels")

	wq := new(sync.WaitGroup)
	wq.Add(2)

	channel := make(chan int)
	defer close(channel)

	go writeChannel(wq, channel, 5)
	go readChannel(wq, channel, 5)
	wq.Wait()
}

func readChannel(wq *sync.WaitGroup, channel chan int, size int) {
	defer wq.Done()
	for i := 0; i < size; i++ {
		fmt.Println(<-channel)
	}
}

func writeChannel(wq *sync.WaitGroup, channel chan int, size int) {
	defer wq.Done()
	for i := 0; i < size; i++ {
		channel <- i
	}
}
