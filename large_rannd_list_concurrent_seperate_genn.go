package main

/*
Write 50billion numbers to a file
*/
import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

var mu sync.Mutex
var wg sync.WaitGroup

func generateRand(ch chan *[10000]int) {

	var temp int
	i := 0
	var result = [10000]int{}

	for i < len(result) {
		temp = rand.Int()
		result[i] = temp
		i++
	}
	ch <- &result
	wg.Done()
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var out strings.Builder
	rand.Seed(time.Now().UnixNano())

	var start = time.Now()
	var ch = make(chan *[10000]int, 100) //channel of size 10
	//var totalGenerated = 0
	var bound, i, threads int = 1000000, 0, 100

	for i < threads {
		wg.Add(1)
		go generateRand(ch)
		i++
	}

	wg.Wait()
	close(ch)
	var end = time.Since(start).Seconds()
	var timeToPrintStart = time.Now()
	k := 0

	for {
		arr, ok := <-ch
		if ok {
			for i :=0; i < len(*arr); i++  {
				out.WriteString("Random value:" + strconv.Itoa(i) + "==>" + strconv.Itoa(arr[i]) + "\n")
				k++
			}
		} else {
			break
		}
	}
	fmt.Print(out.String())
	var timePrintEnd = time.Since(timeToPrintStart).Seconds()

	fmt.Print("\n\nGenerated ", bound, " strings in:", end, "seconds\n")
	fmt.Print("Printed ", bound, " strings in:", timePrintEnd, " seconds\n\n")

}
