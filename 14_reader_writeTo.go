package main

import (
	"bufio"
	"fmt"
	"io"
)

type R struct {
	n int
}
type Writer struct {
}

func (w *Writer) Write(d []byte) (n int, e error) {
	fmt.Printf("%s\n", d)
	return len(d), nil
}

var asciiNNumStart = 97
const noOfLetters int = 25

func (r *R) Read(p []byte) (n int, err error) {

	if r.n == 0 {
		r.n = asciiNNumStart //we start from char(97)
	}
	fmt.Printf("Reading ascii #%d\n", r.n) //show the current ascii value we're reading into p
	if r.n <= asciiNNumStart+noOfLetters {

		copy(p, string(byte(r.n))) //store the current ascii code as a string/rune in p
		r.n += 1                   //go to next ascii number
		return 1, nil
	}
	return 0, io.EOF
}

func main() {
	_R := &R{} //or new(R) which also returns a pointer
	r := bufio.NewReader(_R)

	bw := bufio.NewWriterSize(new(Writer), 3) //internal buffer =3, so flush (write) after every 3 reads

	//this write read and buffers up to the max num of characters needed by the writer (3)

	//n, err := r.WriteTo(bw) or readFrom below are equivalent
	n, err :=bw.ReadFrom(r)
	if err != nil {
		panic(err)
	}
	bw.Flush()
	fmt.Printf("Written bytes: %d\n", n)
}
