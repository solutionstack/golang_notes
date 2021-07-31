package main

import (
	"bufio"
	"fmt"
)
type Writer int

func (*Writer) Write(d []byte) (n int, err error){ //our simple writer object of type Writer
	fmt.Println(len(d))
	return len(d), nil
}
func main(){

	fmt.Println("Starting unbuffered IO")
	w:= new(Writer) //write unbuffered data
	w.Write([]byte{'f'})
	w.Write([]byte{'b'})
	w.Write([]byte{'c'})

	fmt.Println("Starting buffered IO")
	b:= bufio.NewWriterSize(w, 2)//write buffered data out, every 2 calls (i,e its size)
	b.Write([]byte{'e'})
	b.Write([]byte{'f'})
	b.Write([]byte{'g'})
	b.Flush()//flush so pending data on the buffer is written



}
