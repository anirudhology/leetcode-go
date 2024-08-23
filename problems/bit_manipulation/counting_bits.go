package bit_manipulation

func CountingBits(n int) []int {
	// Lookup table to store count of set bits
	lookup := make([]int, n+1)
	// Offset to keep track of powers of 2
	offset := 1
	for i := 1; i <= n; i++ {
		if offset*2 == i {
			offset *= 2
		}
		lookup[i] = 1 + lookup[i-offset]
	}
	return lookup
}
