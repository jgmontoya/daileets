package single_number

import "testing"

func CheckSingleNumber(t *testing.T, input []int, solution int) {
	single_number := singleNumber(input)
	if single_number == solution {
		t.Logf("OK")
	} else {
		t.Errorf("singleNumber(%d). Got %d, expected %d", input, single_number, solution)
		t.Fail()
	}
}

func TestSingleNumber(t *testing.T) {
	input := []int{2, 2, 1}
	solution := 1
	CheckSingleNumber(t, input, solution)

	input = []int{4, 1, 2, 1, 2}
	solution = 4
	CheckSingleNumber(t, input, solution)

	input = []int{1}
	solution = 1
	CheckSingleNumber(t, input, solution)
}
