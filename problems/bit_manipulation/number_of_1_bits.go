package bit_manipulation

func NumberOf1Bits(n int) int {
	setBits := 0
	for n > 0 {
		n &= (n - 1)
		setBits++
	}
	return setBits
}
