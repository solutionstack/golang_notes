package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	dup_lines := make(map[string]int)
	input:= bufio.NewScanner(os.Stdin)

	for input.Scan(){
		dup_lines[input.Text()]++
	}
	for line,n := range dup_lines{
		if n > 1{
			fmt.Printf("%s appears, %d times\n\n", line, n)
		}
	}
}
