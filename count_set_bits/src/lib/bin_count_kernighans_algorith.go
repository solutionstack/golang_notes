package lib

func CountBitsKernighan(v uint64) int {
	/**
	 Kerighann's algorithm uses the following steps.. usually in a loop, that takes the number as a condition
	so it loops while the number is > 0

	*As a rule, when you subtract 1 from a number, the resultant number in binary form, has all the bits after the
	right-most set bit flipped, including the rightmost set bit
	i.e 115 =  01110011
	115-1 = 114 = 01110010 ...so the rightmost set bit has been flipped, so tht bit is counted,
	                       ..then we subtract again and keep counting till no more set bit

	in a program
	LOOP START (while the number > 0)
	On each loop iteration, we are reducing the set bits by one


	first loop <<
	115 = 01110011
	115-1 = 01110010

	01110011 & 01110010 = 01110010 (114)


	2nd loop <<
	114 = 01110010
	114-1 =01110001

	01110010 & 01110001 = 01110000 (112)

	3rd loop <<
	112 = 01110000
	112-1 =01101111

	01110000 & 01101111 = 01100000 (96)

	4th loop <<
	96 =   01100000
	96-1 = 01011111

	01100000 & 01011111 = 01000000 (64)


	5th loop <<
	01000000 (64)
	00111111 (64-1)

	AND =00000000 (0)

	LOOP END


	So we made 5 iterations before getting to zero, meaning the number of set bits was 5

	*/
	var count int
	for v != 0 {
		count += 1
		v = v & (v - 1)
	}

	return count
}
/**

 */
//count holds the 'rest' variables, so is an array
// but in our case we only use the first index to store values
func CountBitsKernighanRecursive(v uint64, count ...int) int {

	if len_ := len(count); v == 0 && len_ == 0 {
		return 0
	} else if v == 0 && len_ > 0 {//no more b its to check
		return count[0]
	}

	if len(count) == 0 {
		count = append(count, 1)
	} else {
		count[0] = count[0] + 1 //increase the count variable on each loop
	}

	//for the next loop, compute the "value AND value-1", to discard the rightmost set bit
	return CountBitsKernighanRecursive(v&(v-1), count...)

}
