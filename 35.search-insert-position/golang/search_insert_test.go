package search_insert

import "testing"

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}

	t.Log("should return the index if the target is present")
	target := 5

	search_insert := SearchInsert(nums, target)
	if search_insert != 2 {
		t.Errorf("Got %d; expected 2", search_insert)
		t.Fail()
	} else {
		t.Log("OK")
	}

	t.Log("should return the index where it should be inserted if not present")
	target = 2

	search_insert = SearchInsert(nums, target)
	if search_insert != 1 {
		t.Errorf("Got %d; expected 1", search_insert)
		t.Fail()
	} else {
		t.Log("OK")
	}

	t.Log("should consider the new index of the bigger array")
	target = 7

	search_insert = SearchInsert(nums, target)
	if search_insert != 4 {
		t.Errorf("Got %d; expected 4", search_insert)
		t.Fail()
	} else {
		t.Log("OK")
	}
}
