package main

import (
	"bufio"
	"fmt"
)

type R struct{}

func (r *R) Read(p []byte) (n int, e error) {
    fmt.Print("Reading\n", )

    copy(p, "abcd")//copies to len(p)
    return len(p), nil
}


func main(){
	r := new(R)
	buf := make([]byte, 2)

	n, err := r.Read(buf)//this is a non buffered read, would read two bytes into the buffer as that's the buffers' capacity
	fmt.Printf("read = %s, n=%d\n", string(buf[:]), n)

	buf = make([]byte,4) //this extends the buffers capacity to 4


	n, err = r.Read(buf) //would read again from the internal buffer to buf's capacity, as we are using a non-buffered Read
    if err  != nil{
    	panic(err)
	}
    fmt.Printf("read = %s, n=%d\n", string(buf[:]), n)


    //now using a buffered read
    br := bufio.NewReader(r)
    br.
	buf = make([]byte,2) //reduce capacity of same buffer to 2
    n, err = br.Read(buf)
	fmt.Printf("read = %s, n=%d\n", string(buf[:]), n)

	buf = make([]byte,4) //increase capacity to four
	n, err = br.Read(buf) //here we read in just two chars from br into buf, since it was previously buffered internally, instead from reading from the underlying
	                      //*R.Read reader
	fmt.Printf("read = %s, n=%d\n", string(buf[:]), n)

}