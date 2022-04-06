package single_number

func singleNumber(nums []int) int {
	result := nums[0]
	for index := 1; index < len(nums); index++ {
		result ^= nums[index]
	}
	return result
}
