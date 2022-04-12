package merge

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1_pointer := m - 1
	nums2_pointer := n - 1

	for index := n + m - 1; index >= 0; index-- {
		if nums1_pointer >= 0 && nums2_pointer >= 0 {
			if nums1[nums1_pointer] >= nums2[nums2_pointer] {
				nums1[index] = nums1[nums1_pointer]
				nums1_pointer--
			} else {
				nums1[index] = nums2[nums2_pointer]
				nums2_pointer--
			}
		} else if nums1_pointer >= 0 {
			nums1[index] = nums1[nums1_pointer]
			nums1_pointer--
		} else {
			nums1[index] = nums2[nums2_pointer]
			nums2_pointer--
		}
	}
}
