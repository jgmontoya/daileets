package intersect

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

func CheckIntersect(t *testing.T, nums1 []int, nums2 []int, solution []int) {
	intersection := intersect(nums1, nums2)
	if SliceEqual(intersection, solution) {
		t.Log("OK")
	} else {
		t.Errorf("interserct([%d], [%d]): Got %d; expected %d", nums1, nums2, intersection, solution)
		t.Fail()
	}
}

func TestIntersect(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	solution := []int{2, 2}
	CheckIntersect(t, nums1, nums2, solution)

	nums1 = []int{4, 9, 5}
	nums2 = []int{9, 4, 9, 8, 4}
	solution = []int{9, 4}
	CheckIntersect(t, nums1, nums2, solution)
}
