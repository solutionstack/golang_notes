package main


import (
	"fmt"
	"musings/count_set_bits/src/lib"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	for _, v := range args {
		if n, err := strconv.Atoi(v); err == nil {
			//ret := lib.CountBitsKernighanRecursive(uint64(n))
			ret := lib.CountBits(uint64(n))

			fmt.Printf("Number of set bits in [%s] is %d\n", v, ret)
		} else {
			fmt.Fprintf(os.Stderr, "Error in input argument: '%s' is invalid\n\n", v)
			//should we return on input error
		}
	}
}
