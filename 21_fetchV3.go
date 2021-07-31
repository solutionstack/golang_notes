package main

///
/// Fetch urls sent as program arguments and dump their content to stdOut
///
import (
	//"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

//type Writer struct {
//}
//
//func (*Writer) Write(d []byte) (n int, err error) {
//	//fmt.Println(string(d))
//
//	return len(d), nil
//}
func main() {

	start := time.Now()
	for _, url := range os.Args[1:] {

		reqStartAt := time.Now()
		_, err := http.Get(url)

		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		//w := new(Writer)
		//bw := bufio.NewWriterSize(w, 64)

		//fmt.Println("\n\nHeaders:", resp.Header)
		//fmt.Println("\nStatus:", resp.Status)
		//fmt.Println("\nStatusCode:", resp.StatusCode, "\n\n")

		timeCurrentReq := time.Since(reqStartAt).Seconds()
		fmt.Printf("\nRequest for %s took: %f\n", url, timeCurrentReq)

		//_, err = io.Copy(bw, resp.Body) //copy till EOF

		//if err != nil && err != io.EOF {
		//
		//	fmt.Fprint(os.Stderr, err)
		//	os.Exit(1)
		//}
	}

	end := time.Since(start).Seconds()

	fmt.Printf("\nTotal Duration: %f\n\n", end)

}
