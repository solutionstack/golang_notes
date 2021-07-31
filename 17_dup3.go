package main

///
///  PROGRAM TO PRINT OUT DUPLICATE LINES IN A FILE, SHOWING EACH LINE NUMBER CONTAINING DUPLICATES IN EACH FILE
///
import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	//"strings"
)

func main() {

	files := os.Args[1:]
	resultMap := map[string][]int{} //initialize an empty map, th key is a line in the file,
	//while the value ([]int) holds the the line numbers where a line is duplicated

	for _, arg := range files {
		file, err := os.Open(arg)
		if err != nil {
			panic(err)
		}
		countLines2(file, resultMap)

		////print the duplicate line for each file
		fmt.Println("\nFile: >>>", file.Name(), "; Duplicates")
		printDuplicates(resultMap)
		resultMap = map[string][]int{} //reset the m
	}
	fmt.Println()
}

func printDuplicates(m map[string][] int) {

	for line, lineCount := range m {
		if len(lineCount) > 1 { //if the array has a size greater than 1
		mapJson, err := json.Marshal(lineCount) //convert the []int to json
		if err != nil {
			panic(err)
		}
		fmt.Println("Text -", line, "- duplicated on lines ", string(mapJson)) //marshall above returns a []byte
		//so convert to []string
		}

	}

}

func countLines2(f *os.File, duplicateMap map[string][]int) {
	r := bufio.NewReader(f)
	var line []byte
	var err error
	atEof := false
	lineNumber := 0 //keep count of the line number we are on

	for !atEof {
		line, _, err = r.ReadLine() //discard new line \n
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			atEof = true //would cause to break
		}
		lineNumber++

		if string(line) == "" { //skip empty lines
			continue
		}

		//duplicateMap[string(line)] points to the array for that index in the map
		//so on each loop re assign the array..
		//by appending the the new value (lineNumber) to a copy of the array at that position

		duplicateMap[string(line)] = append(duplicateMap[string(line)], lineNumber)

		/*
		  alternate way to populate map below
		*/
		//if len(duplicateMap[string(line)]) == 0 { //this line hasn't been previously stored in a map
		//	duplicateMap[string(line)] = []int{lineNumber}
		//} else {
		//	//append the new line number to the duplicate array for the current line of text
		//	duplicateMap[string(line)] = append(duplicateMap[string(line)], lineNumber)
		//}

	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
}
