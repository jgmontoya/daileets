package minimum_total

import "testing"

func CheckMinimumTotal(t *testing.T, input [][]int, solution int) {
	minimum_total := minimumTotal(input)
	if minimum_total == solution {
		t.Logf("OK")
	} else {
		t.Errorf("Got %d; expected %d", minimum_total, solution)
		t.Fail()
	}
}
func TestMinimumTotal(t *testing.T) {
	input := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	solution := 11
	CheckMinimumTotal(t, input, solution)

	input = [][]int{{-10}}
	solution = -10
	CheckMinimumTotal(t, input, solution)
}
