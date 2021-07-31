package main

import "fmt"

/**
Reverse an array or slice
*/

func reverse(slice []string) {
	//make a copy of the original slice, as we would invalidate the slice if we read/write into it simultaneously
	tmp := make([]string, cap(slice))
	copy(tmp, slice)
	for i, _ := range tmp {
		slice[(cap(slice)-i)-1] = tmp[i]
	}
}
type Point struct {
	X,Y int64
}
func (p Point) Distance(x,y int64) int64{
	return x+y
}
func (p *Point) ScaleBy(pt *Point,y int64) int64{
	return y
}
func main() {
	str := []string{"a", "b", "u"}


	reverse(str[:])
	fmt.Printf("\n%v\n", str)



	p := Point{1, 2}


	distance := p.Distance // value expression fmt.Println(distance(x, y)) // "5" fmt.Printf("%T\n", distance) // "func(Point, Point) float64"

	distance(1,2)
	scale := (*      ).ScaleBy // method expresiion "{2 4}" fmt.Printf("%T\n", scale) // "func(*Point, float64)"

scale(&p,  &p, int64(2))
}
