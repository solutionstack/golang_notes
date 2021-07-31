package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func main() {
	s := strings.NewReader(strings.Repeat("b", 60)) //all b's
	br := bufio.NewReader(s) //buffered reader slice of strings


	buf := &bytes.Buffer{}//buffer with Reader/Writer Methods, alternatively .. new(bytes.Buffer)
[]int16
	//bw := bufio.NewWriter(buf) //buffered writer
	bw := bufio.NewWriterSize(buf, 2) //buffered writer
	rw := bufio.NewReadWriter(br, bw)// R/W

	buf2 := make([] byte, 1) //some buffer
	_, err := rw.Read(buf2) //read len(buf2) amount of data into buf2 from br

	if err != nil {
		panic(err)
	}
	fmt.Println(buf2) //should give 98 --> byte (ascii) position of B

	buf2 = []byte("hi") //another random slice
	_, err = rw.Write(buf2) //write that slice to the buffered-writer --> bw, the internal buffer takes 2ß
	                        // chars so it would be filled up, but would not flush


	if err != nil {
		panic(err)
	}


	fmt.Println(buf.String()) //nothing written, bw buffers till its size
	err = rw.Flush() //force write to buf
	fmt.Println(buf.String())

	//err = rw.Flush() // we flush because the last-write would buffer internally (till it reaches it's default length
	if err != nil {
		panic(err)
	}


}
