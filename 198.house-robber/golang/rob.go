package rob

func rob(nums []int) int {
	return memoized_rob(nums, map[int]int{})
}

func memoized_rob(nums []int, memory map[int]int) int {
	if val, ok := memory[len(nums)-1]; ok {
		return val
	}
	if len(nums) <= 2 {
		memory[len(nums)-1] = sliceMax(nums)
		return memory[len(nums)-1]
	}

	last_element := nums[len(nums)-1]
	nums_without_last := nums[:len(nums)-1]
	nums_wihtout_last_two := nums[:len(nums)-2]

	options := []int{memoized_rob(nums_without_last, memory), memoized_rob(nums_wihtout_last_two, memory) + last_element}

	memory[len(nums)-1] = sliceMax(options)
	return memory[len(nums)-1]
}

func sliceMax(nums []int) int {
	max := nums[0]
	for _, elem := range nums {
		if elem > max {
			max = elem
		}
	}
	return max
}
