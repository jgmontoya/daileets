package search

import "testing"

func TestSearch(t *testing.T) {
	t.Log("Should return -1 when target is not present")
	input_array := []int{-1, 0, 3, 5, 9, 12}

	search_result := Search(input_array, 2)
	if search_result != -1 {
		t.Errorf("Got %d; expected -1", search_result)
		t.Fail()
	} else {
		t.Log("OK")
	}

	t.Log("Should return the index when target is present")
	search_result = Search(input_array, 9)
	if search_result != 4 {
		t.Errorf("Got %d; expected 4", search_result)
		t.Fail()
	} else {
		t.Log("OK")
	}

	search_result = Search(input_array, 0)
	if search_result != 1 {
		t.Errorf("Got %d; expected 1", search_result)
		t.Fail()
	} else {
		t.Log("OK")
	}
}
