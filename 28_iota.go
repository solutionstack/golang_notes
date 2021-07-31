package main

import (
	"fmt"
)

const (
	flagUp = 1 << iota //starts enumerating from 0
	flagDown
	flagSetBroadcast
	flagPointToPoint
	flagMulticast
)

func main() {

	fmt.Printf("%d; %b\n", flagUp, flagUp)
	fmt.Printf("%d; %b\n", flagDown, flagDown)
	fmt.Printf("%d; %b\n", flagSetBroadcast, flagSetBroadcast)
	fmt.Printf("%d; %b\n", flagPointToPoint, flagPointToPoint)
	fmt.Printf("%d; %b\n", flagMulticast, flagMulticast)

	fmt.Println(isUp(0))

	val := uint32(0)
	val2 := uint32(0)
	val3 := uint32(0)

	fmt.Println(val) //0
	turnUp(&val)
	fmt.Println(val) //1

	fmt.Printf("original : %d \t %b\n ", val2, val2)
	setMultipleFlags(&val2, flagMulticast, flagPointToPoint)
	fmt.Printf(" set: %d \t %b\n", val2, val2)

	fmt.Printf("\ntestinng flagUp, flagMulticast in val3 : %v\n", testMultiple(&val3, flagUp, flagMulticast))
	fmt.Printf("\ntestinng flagMulticast, flagPointToPoint in val2 : %v\n", testMultiple(&val2, flagMulticast, flagPointToPoint))
	fmt.Printf("\ntestinng flagMulticast, flagUp in val2 : %v\n", testMultiple(&val2, flagMulticast, flagUp))

	clearMultiple(&val2, flagMulticast)
	fmt.Printf(" set: %d \t %b\n", val2, val2)

	clearMultiple(&val2, flagPointToPoint)
	fmt.Printf(" set: %d \t %b\n", val2, val2)

	/**
	  USe iota to generate byte sizes
	*/

	const (
		_= 1 << (10 * iota) //1 << (10 * 0) == 0
		KiB                   //1 << (10 * 1) === 1024, each shift is multiply by 2, so 2^10 === 1024
		MiB
		GiB
		TiB
		PiB
		EiB
		ZiB //need external packages to hold numbers > uint64
		YiB
	)
	fmt.Println(KiB)
	fmt.Println(MiB)
	fmt.Println(GiB)
	fmt.Println(TiB)
	fmt.Println(PiB)
	fmt.Println(EiB)

}

func isUp(v int) bool {
	return v&flagUp == flagUp //if the flagUp bit is set in V
}
func turnUp(v *uint32) {
	*v |= flagUp
}

func setMultipleFlags(v *uint32, flags ...int) {
	for _, flag := range flags {
		*v |= uint32(flag)
	}
}

func testMultiple(v *uint32, flags ...int) bool {
	for _, flag := range flags {
		if uint32(flag) == (*v & uint32(flag)) { //test each bit (or v & (flag1 | flag2 |... | flagN)
		} else {
			return false //bit wasnt set for some flags
		}
	}
	return true
}

func clearMultiple(v *uint32, flags ...int) {
	for _, flag := range flags {
		*v &= ^uint32(flag) //or v & ^(flag1 | flag2 | ...)
	}
}
