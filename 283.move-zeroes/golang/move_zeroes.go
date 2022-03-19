package move_zeroes

func MoveZeroes(nums []int) {
	index := 0
	next_non_zero := NextNonZero(nums, index)
	for next_non_zero != -1 {
		swap(nums, index, next_non_zero)
		index++
		next_non_zero = NextNonZero(nums, index)
	}
}

func NextNonZero(nums []int, from int) int {
	for index := from; index < len(nums); index++ {
		if nums[index] != 0 {
			return index
		}
	}
	return -1
}

func swap(nums []int, first int, second int) {
	mem := nums[first]
	nums[first] = nums[second]
	nums[second] = mem
}
