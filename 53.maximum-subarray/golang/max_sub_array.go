package max_sub_array

func maxSubArray(nums []int) int {

	local_max := 0
	global_max := nums[0]

	for index := 0; index < len(nums); index++ {
		local_max = max(nums[index], nums[index]+local_max)
		if local_max > global_max {
			global_max = local_max
		}
	}

	return global_max
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
