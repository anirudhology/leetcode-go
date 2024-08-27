package math

func HappyNumber(n int) bool {
	// Special case
	if n == 1 || n == 7 {
		return true
	}
	for n > 9 {
		n = calculateSquareOfDigits(n)
		if n == 1 || n == 7 {
			return true
		}
	}
	return false
}

func calculateSquareOfDigits(n int) int {
	square := 0
	for n > 0 {
		remainder := n % 10
		square = square + remainder*remainder
		n = int(n / 10)
	}
	return square
}
