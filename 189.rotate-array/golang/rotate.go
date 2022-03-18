package rotate

func Rotate(nums []int, k int) {
	nums_length := len(nums)
	real_rotation := k % nums_length
	break_point := nums_length - real_rotation

	rotated := append(nums[break_point:], nums[:break_point]...)
	for index := 0; index < nums_length; index++ {
		nums[index] = rotated[index]
	}
}
