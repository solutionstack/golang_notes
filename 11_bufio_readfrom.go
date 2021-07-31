package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Writer4 struct {
}

func (*Writer4) Write(d []byte) (n int, err error) { //our simple writer object of type Writer
	fmt.Println(string(d))
	return len(d), nil
}

func main() {
	s := strings.NewReader("foobarr")
	w := new(Writer4)

	bufferWrite := bufio.NewWriterSize(w, 3)

	bufferWrite.ReadFrom(s) //use s as our source, ReadFrom reads all data from the source and writes them to the buffer

	err := bufferWrite.Flush() //flush in case data is left in buffer
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}


	///buffered Reader
	l := strings.NewReader("abcdefghijklmnopqrstuv")
	r := bufio.NewReader(l)
	bt := make([]byte, 3)

	for {
		n, err := r.Read(bt)
		fmt.Println(r.Peek(n)) //peek the next n bytes
		fmt.Println(n, err, bt[:n], string(bt[:n]))

		if err == io.EOF {
			break
		}
	}

	/*s1 := strings.NewReader(strings.Repeat("a", 16) + strings.Repeat("b", 16))
	r := bufio.NewReaderSize(s1, 16)
	b, _ := r.Peek(3)
	fmt.Println(s1)
	fmt.Printf("%q\n", b)
	r.Read(make([]byte, 16))
	r.Read(make([]byte, 15))
	fmt.Printf("%q\n", b)*/
}
