package main

import (
	"fmt"
	"os"
)

func main() {
	s := ""
	sep:= "\n"

	for key, arg := range os.Args[1:] {
		s += fmt.Sprintf("%d => %s %s",key, arg , sep)
	}
	fmt.Println(s)
}
