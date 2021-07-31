package main

import (
	"fmt"
	"strings"
)

func eliminateAdjacentDuplicate(slice *[]string) {

	result := make ([] string, len(*slice))
	g:=0 //keeps count of our result modifications

	//loop through the pointed slice, and add the zeroth index to the result slice
	for i, v := range *slice{
		if i == 0{
			result[g] = v
		}

		if i > 0{
			//skip duplicated indexes (i.e if the current char, equals the last
			if v == (*slice)[i-1]{
				continue
			}
			result[g] = v

		}
		g++
	}
	result = result[:g] //resize the slice to reflect just the elements that whhere preserved
	*slice = result

}

func main() {
	str := "aaabbbcccdefggghiii"

	slice_ := strings.Split(str, "")
	eliminateAdjacentDuplicate(&slice_)

	fmt.Printf("\n%v\n", slice_)//print the modified slice


}
