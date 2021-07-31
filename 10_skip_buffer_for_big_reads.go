package main

import (
	"bufio"
	"fmt"
	"strings"
)

type R2 struct {
}

func (*R2) Read(p []byte) (n int, e error) {
	fmt.Println("Reading...")
	s := strings.Repeat("b", len(p))
	copy(p, s)
	return len(p), nil
}

func main() {
	r := new(R2)
	br := bufio.NewReaderSize(r, 10) //internal buffer can hold max 10 chars
	bt := make([]byte, 8)

	fmt.Println()
	n, err := br.Read(bt) //attempt to buffer 17 bytes into internal buffer of size 10, s
	// kips buffering, and reads directly into byte.slice

	if err != nil {
		panic(err)
	}

	fmt.Printf("Read:%s\n Read-count:%d\n Buffered-internally:%d ", string(bt[:]), n, br.Buffered())

}
