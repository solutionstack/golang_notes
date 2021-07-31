package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(intToString([]int{1, 2, 3})) // "[1, 2, 3]" }
}
func intToString(values []int) string {

	b := &bytes.Buffer{}
	(b).WriteByte('[')

	for i, v := range values {

		if i > 0 {
			err := b.WriteByte(',')
			if err != nil {
				panic(err)
			}
		}
		//we use printf instead of WriteByte|Rune as the numeric would be converted to their ascii equivalents
		//so 1,2,3 would likely be unprintable control characters
		fmt.Fprintf(b, "%d", v)
	}
	b.WriteByte(']')
	return string(b.String())

}
