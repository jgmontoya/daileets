package permute

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	permutations := [][]int{}

	for index, num := range nums {
		for _, subpermutation := range permute(remove(nums, index)) {
			permutations = append(permutations, insertNumAtStart(num, subpermutation))
		}
	}
	return permutations
}

func insertNumAtStart(num int, subpermutation []int) []int {
	result := make([]int, len(subpermutation)+1)
	result[0] = num
	for index, sub_num := range subpermutation {
		result[index+1] = sub_num
	}
	return result
}

func remove(nums []int, index_to_remove int) []int {
	result := make([]int, len(nums)-1)
	offset := 0
	for index, num := range nums {
		if index == index_to_remove {
			offset = 1
			continue
		}
		result[index-offset] = num
	}
	return result
}
