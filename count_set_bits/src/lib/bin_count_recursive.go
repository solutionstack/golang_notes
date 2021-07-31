package lib

func CountBitsRecursive(v uint64) uint8{

	if v == 0{
		return uint8(v)
	}  else{
		return uint8(v & 1) + uint8(CountBitsRecursive(v>>1))
	}
}
