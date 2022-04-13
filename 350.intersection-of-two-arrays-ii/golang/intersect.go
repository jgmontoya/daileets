package intersect

func intersect(nums1 []int, nums2 []int) []int {
	dict := map[int]int{}
	result := []int{}

	for _, num := range nums1 {
		dict[num]++
	}

	for _, num := range nums2 {
		if dict[num] > 0 {
			result = append(result, num)
			dict[num]--
		}
	}

	return result
}
