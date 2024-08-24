package math

func MissingNumber(nums []int) int {
	n, sum := len(nums), 0
	for _, num := range nums {
		sum += num
	}
	return (n * (n + 1) / 2) - sum
}
