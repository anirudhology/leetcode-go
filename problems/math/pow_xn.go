package math

func PowXN(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / x * PowXN(1/x, -n-1)
	}
	if n%2 == 0 {
		return PowXN(x*x, int(n/2))
	}
	return x * PowXN(x*x, int(n/2))
}
