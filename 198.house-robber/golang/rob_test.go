package rob

import "testing"

func CheckRob(t *testing.T, input []int, solution int) {
	robbery := rob(input)
	if robbery == solution {
		t.Logf("OK")
	} else {
		t.Errorf("Got %d; expected %d", robbery, solution)
		t.Fail()
	}
}
func TestRob(t *testing.T) {
	input := []int{1, 2, 3, 1}
	solution := 4
	CheckRob(t, input, solution)

	input = []int{2, 7, 9, 3, 1}
	solution = 12
	CheckRob(t, input, solution)

	input = []int{2, 1, 1, 2}
	solution = 4
	CheckRob(t, input, solution)
}
