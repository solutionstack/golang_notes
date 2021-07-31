package main
import (
"fmt"
"sync"
)
var x  = 0
var mu = sync.Mutex{}
func increment(wg *sync.WaitGroup) {
	mu.Lock()
	defer mu.Unlock()
	x = x + 1

	wg.Done()
}
func main() {
	var w sync.WaitGroup
	var j int
	for i:=0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)

		j = i
	}
	w.Wait()
	fmt.Printf("final value of x=%d, i=%d", x, j)
}
