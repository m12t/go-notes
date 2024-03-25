package main

import (
	"fmt"
)

type ByteCounter int
type ByteMux int

func (c *ByteCounter) Write(p []byte) (int, error) {
	fmt.Println("test")
	return len(p), nil
}

func main() {
	var (
		b ByteCounter
		c ByteMux // doesn't implement io.Writer
	)
	b = 23

	name := "m"
	fmt.Fprintf(&b, "hello, %s", name)
	// fmt.Fprintf(&c, "hello, %s", name)  // err
}
