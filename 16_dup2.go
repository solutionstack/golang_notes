package main

/**
 * Print out duplicate lines from files
 */
import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	countsMap := make(map[string]int)

	files := os.Args[1:]

	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file \n %v", err)
			os.Exit(-1)
		}
		countLines(f, countsMap)
	}

	for text, count := range countsMap {
		if count > 1{
			fmt.Println("(",text,") occur's " , count, " times")
		}
	}
}

func countLines(f *os.File, countsMap map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		countsMap[input.Text()]++ //increase the value of each line the text is repeated
		                          //on an initial entry countsMap[input.Text()] would be set to 0
	}
	if err := input.Err(); err != nil {
		panic(err)
	}

}
