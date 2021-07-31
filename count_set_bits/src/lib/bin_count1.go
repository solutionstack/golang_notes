package lib

func CountBits(v uint64) uint8{

	var count uint8
	for v !=0{
		count += uint8(v & 1)
		v >>= 1
	}
	return count
}
