package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	randomString := "loemipsum ittariun"
	scanner := bufio.NewScanner(strings.NewReader(randomString))

	scanner.Split(bufio.ScanRunes)//split by each utf8 byte
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Bytes())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Bloop!")
	}
}
