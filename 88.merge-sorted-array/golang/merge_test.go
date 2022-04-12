package merge

import "testing"

func SliceEqual(first_slice []int, second_slice []int) bool {
	first_length := len(first_slice)
	second_length := len(second_slice)

	if first_length != second_length {
		return false
	}

	for i := 0; i < first_length; i++ {
		if first_slice[i] != second_slice[i] {
			return false
		}
	}
	return true
}

func CheckMerge(t *testing.T, nums1 []int, m int, nums2 []int, n int, solution []int) {
	t.Logf("merge(%d, %d, %d, %d):\n", nums1, m, nums2, n)
	merge(nums1, m, nums2, n)
	if SliceEqual(nums1, solution) {
		t.Log("OK")
	} else {
		t.Errorf("Got %d, expected %d", nums1, solution)
		t.Fail()
	}
}

func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	solution := []int{1, 2, 2, 3, 5, 6}
	CheckMerge(t, nums1, m, nums2, n, solution)

	nums1 = []int{1}
	m = 1
	nums2 = []int{}
	n = 0
	solution = []int{1}
	CheckMerge(t, nums1, m, nums2, n, solution)

	nums1 = []int{0}
	m = 0
	nums2 = []int{1}
	n = 1
	solution = []int{1}
	CheckMerge(t, nums1, m, nums2, n, solution)

	nums1 = []int{4, 5, 6, 0, 0, 0}
	m = 3
	nums2 = []int{1, 2, 3}
	n = 3
	solution = []int{1, 2, 3, 4, 5, 6}
	CheckMerge(t, nums1, m, nums2, n, solution)
}
