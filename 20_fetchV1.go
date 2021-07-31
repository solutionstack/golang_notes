package main
///
/// Fetch urls sent as program arguments and dump their content to stdOut
///
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Println(string(body))
	}

}
