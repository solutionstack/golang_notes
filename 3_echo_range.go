package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	sep:= ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "

	}
	fmt.Println(strings.Join(os.Args[1:],","))//one liner
	fmt.Println(s)
}
