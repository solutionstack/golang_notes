package main
//rest a buffer to be used in different writers
import (
	"bufio"
	"fmt"
)

type Writer2 int

type Writer3 int //writer1 declared in diff file

func (*Writer2) Write(d []byte) (n int, err error) { //our simple writer object of type Writer
	fmt.Println(d)
	return len(d), nil
}
func (*Writer3) Write(d []byte) (n int, err error) { //our simple writer object of type Writer
	fmt.Println(string(d))
	return len(d), nil
}
func main() {

	fmt.Println("Starting buffered IO")
	w := new(Writer2)  //1st writer
	w2 := new(Writer3) //...

	buffW := bufio.NewWriterSize(w, 3)
	buffW.Write([]byte{'b'})

	//reset for a new writer
	buffW.Reset(w2)



	//even though our buffer is meant to hold just three chars, this call would print all data
	//as the writer would detect that the data is too large to fit in the buffer
	buffW.Write([]byte("cooo"))

	//lets use the buffering by writing data in chunks
	buffW.Write([]byte("aa")) //this shouldn't be printed as the we haven't reached the buffer's capacity (3)
	buffW.Write([]byte("b"))  //we would reach the buffers capacity here, but still would'nt print


	buffW.Write([]byte("c")) //now this would would try to go over the buffers cap. so the prev. data gets flushed
	// we therefore get "aab" written out (flushed), and 'c' as the only data

	//let see what free space we have at this point
	fmt.Println("Available spcae:", buffW.Available()) //should be 2, as we have only 'c' in the buffer

	//finally flush the remaining data
	buffW.Flush()

}
