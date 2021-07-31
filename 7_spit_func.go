
//A Program to count duplicate words, separated by #

package main

import (
	"bufio"
	"fmt"
	"os"
)

func dropCR(data []byte) []byte {
	//ignore, new line characters... Would this work on windows, which uses \n\r??
	if len(data) > 0 && (data[len(data)-1] == '\r' || data[len(data)-1] == '\n') {
		return data[0 : len(data)-1]//return all but the last char, as slices in go, sre exclusive of the last range
	}
	return data
}

//split function to detect the separator '#' and return the number of tokens to consume
func split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i := 0; i < len(data); i++ {
		if data[i] == '#' {
			return i + 1, dropCR(data[:i]), nil
		}
	}
	if !atEOF {
		return 0, nil, nil
	}
	return len(data), dropCR(data), bufio.ErrFinalToken
}

func main() {
	dup_lines := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	input.Split(split)
	for input.Scan() {
		dup_lines[input.Text()]++
	}
	for line, n := range dup_lines {
		if n > 1 {
			fmt.Printf("%s appears, %d times\n\n", line, n)
		}
	}
	if err:= input.Err(); err != nil{
		fmt.Fprintln(os.Stderr," Error reading STDIN", err)
	}
}
