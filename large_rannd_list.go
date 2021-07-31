package main


/*
Write 50billion numbers to a file
 */
import (
	"fmt"
	"math/rand"
	"time"
)

func generateRand(intArr *[]int, generatedCount *int) {

}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	var start = time.Now()
	//var totalGenerated = 0
	var bound, i int = 1000000, 0

	//var buffer = &bytes.Buffer{}

	//f, err := os.Create("number.txt")
	//check(err)
	//var w = bufio.NewWriterSize(f, 1024*10)


	//defer f.Close()

	for i < bound {
		temp := rand.Int()

		//_, err = w.WriteString(fmt.Sprintf("%d\n", temp))
		//check(err)
		fmt.Printf("generating number => %d \t",i )
		fmt.Println(temp)
		i++
	}

	var end = time.Since(start).Seconds()
	fmt.Print("wrote ", bound, " strings in:", end, "seconds")

}
