package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Working with interfaces")
	c := &Capper{os.Stdout}
	fmt.Fprintln(c, "Ssup file writer?")
}

// This implements the io.Writer inteface. So it must implemet the Write method
type Capper struct {
	wtr io.Writer
}

func (c Capper) Write(b []byte) (int, error) {
	out := make([]byte, len(b))
	diff := byte('a' - 'A')

	for i, c := range b {
		if c >= 'a' && 'c' <= 'z' {
			c -= diff
		}
		out[i] = c
	}
	return c.wtr.Write(out)
}
