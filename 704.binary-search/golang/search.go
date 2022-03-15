package search

func Search(nums []int, target int) int {
	return RecursiveSearch(nums, target, 0)
}

func RecursiveSearch(nums []int, target int, accumulated_index int) int {
	nums_length := len(nums)
	if nums_length == 0 {
		return -1
	}
	mid_index := nums_length / 2
	mid_element := nums[mid_index]

	if target == mid_element {
		return mid_index + accumulated_index
	} else if target < mid_element {
		return RecursiveSearch(nums[:mid_index], target, accumulated_index)
	} else {
		return RecursiveSearch(nums[mid_index+1:], target, accumulated_index+mid_index+1)
	}
}
