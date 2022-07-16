// Following this blog for channels in Go: https://blog.logrocket.com/concurrency-patterns-golang-waitgroups-goroutines/

package main

import (
	"fmt"
	"sync"
)

func BuffchannelMain() {
	fmt.Println("Printing using buffered channels")

	wq := new(sync.WaitGroup)
	wq.Add(1)

	channel := make(chan int, 2)
	defer close(channel)

	go writeChannel(wq, channel, 2)
	wq.Wait()

	fmt.Println("Contents of the buff channel: ")
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
