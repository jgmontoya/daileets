package search_insert

func SearchInsert(nums []int, target int) int {
	return RecursiveSearchInsert(nums, target, 0)
}

func RecursiveSearchInsert(nums []int, target int, offset int) int {
	nums_length := len(nums)
	if nums_length == 0 {
		return offset
	}

	mid_index := nums_length / 2
	mid_element := nums[mid_index]

	if target == mid_element {
		return mid_index + offset
	} else if target < mid_element {
		return RecursiveSearchInsert(nums[:mid_index], target, offset)
	} else {
		return RecursiveSearchInsert(nums[mid_index+1:], target, offset+mid_index+1)
	}
}
