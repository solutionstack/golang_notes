package main

///
/// Fetch urls sent as program arguments and dump their content to stdOut, using buffer readers/writers to handle the respinse
///
import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)

		fmt.Println(url)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		r := bufio.NewReader(resp.Body)
		w := bufio.NewWriter(os.Stdout)


		_, err = r.WriteTo(w) //write data to stdout as data becomes available. WriteTo calls Read repeatedly to fill
		                     //its buffer

		if err != nil && err != io.EOF {

			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
	}

}
