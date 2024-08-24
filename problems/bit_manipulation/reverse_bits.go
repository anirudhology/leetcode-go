package bit_manipulation

func ReverseBits(n uint32) uint32 {
	var output uint32
	for i := 0; i < 32; i++ {
		output <<= 1
		output += (n & 1)
		n >>= 1
	}
	return output
}
