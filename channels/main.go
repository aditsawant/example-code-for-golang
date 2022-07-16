// Following this blog for channels in Go: https://blog.logrocket.com/concurrency-patterns-golang-waitgroups-goroutines/

package main

import "fmt"

func main() {
	ChannelMain()
	fmt.Println("---------------------------")
	BuffchannelMain()
}
