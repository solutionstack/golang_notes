package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {

	sr := strings.NewReader("aeio.u")
	bf := bufio.NewReader(sr)

	token, err := bf.ReadSlice('.')
	//bf.Read(make([] byte, 3)) //this would invalidate the slice above
	if err != nil && err != io.EOF {
		panic(err)
	}

	fmt.Printf("Token: %q  ", token)

}
