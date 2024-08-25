package bit_manipulation

func SumOfTwoIntegers(a int, b int) int {
	if a == 0 || b == 0 {
		return a | b
	}
	return SumOfTwoIntegers(a^b, (a&b)<<1)
}
