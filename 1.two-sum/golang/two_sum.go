package two_sum

func twoSum(nums []int, target int) []int {
	dict := map[int]int{}

	for index, num := range nums {
		target_partner := target - num

		if val, ok := dict[target_partner]; ok {
			return []int{val, index}
		}
		dict[num] = index
	}
	return []int{0, 1}
}
