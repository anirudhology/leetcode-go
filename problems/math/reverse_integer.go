package math

func ReverseInteger(x int) int {
	maxValue := 2147483647
	num := Abs(x)
	xReverse := 0
	for num != 0 {
		remainder := num % 10
		// Check for overflow
		if (xReverse > (maxValue-remainder)/10) || (xReverse == int(maxValue/10) && remainder > maxValue%10) {
			return 0
		}
		xReverse = xReverse*10 + remainder
		num = int(num / 10)
	}
	if x < 0 {
		return -xReverse
	}
	return xReverse
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
