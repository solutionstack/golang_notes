package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var period = flag.Duration("period", 1*time.Second, "time to sleep")
	flag.Parse()
	fmt.Printf(" Sleeping fo %v...\n", *period)
 	time.Sleep(*period)
	fmt.Println(" exit")

}