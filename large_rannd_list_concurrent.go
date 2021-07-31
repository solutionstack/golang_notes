package main

/*
Write 50billion numbers to a file
*/
import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"sync"
	"time"
)

var mu sync.Mutex
var wg sync.WaitGroup

func generateRand(upperBound *int, counter *int ) {

	temp := rand.Int()

	//_, err = w.WriteString(fmt.Sprintf("%d\n", temp))
	//check(err)
	for *counter < *upperBound {


		mu.Lock()
		*counter++
		mu.Unlock()
		//fmt.Printf("generating number => %d \t: %d\n", *counter, temp)
fmt.Fprint(ioutil.Discard, temp)

	}
	wg.Done()
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	//ch := make(chan string)
	var start = time.Now()
	//var totalGenerated = 0
	var bound, i, threads int = 1000000, 0, 1000
	for i < threads {
		wg.Add(1)
		//go generateRand(&bound, &i, ch)
		go generateRand(&bound, &i)
	}

	wg.Wait()
	var end = time.Since(start).Seconds()
	//close(ch)
	//for output:= range ch {
	//	fmt.Println(output)
	//}

	fmt.Print("wrote ", bound, " strings in:", end, "seconds")

}
