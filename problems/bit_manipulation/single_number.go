package bit_manipulation

func SingleNumber(nums []int) int {
	answer := nums[0]
	for i := 1; i < len(nums); i++ {
		answer ^= nums[i]
	}
	return answer
}
