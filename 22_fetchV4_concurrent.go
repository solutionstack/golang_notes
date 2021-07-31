package main

/**
 * concurrently measure time taken to call and bytes received from arbitrary urls
 */
import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {

	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {

		go fetch(url, ch)
	}
	for i := 0; i < len(os.Args[1:]); i++ {
		fmt.Println(<-ch) //receive from channel
	}
	fmt.Printf("\nTotal time elapsed: %.2fs \n", time.Since(start).Seconds());

}

func fetch(url string, ch chan<- string) {

	start := time.Now()
	res, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint("fetch error: ", err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, res.Body)
	defer res.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("Erro while reading response %s %v", url, err)
		return
	}

	timeTaken := time.Since(start).Seconds()
	ch <- fmt.Sprintf("\nrespnse stats:  time: %.2fs bytes:%d url %s \n", timeTaken, nbytes, url)
}
