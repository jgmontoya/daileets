package oranges_rotting

import "testing"

func CheckOrangesRotting(t *testing.T, input [][]int, solution int) {
	oranges_rotting := OrangesRotting(input)
	if oranges_rotting == solution {
		t.Log("OK")
	} else {
		t.Errorf("Got %d; expected %d", oranges_rotting, solution)
		t.Fail()
	}
}

func TestOrangesRotting(t *testing.T) {
	input := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
	solution := 4
	CheckOrangesRotting(t, input, solution)

	input = [][]int{
		{2, 1, 1},
		{0, 1, 1},
		{1, 0, 1},
	}
	solution = -1
	CheckOrangesRotting(t, input, solution)

	input = [][]int{{0, 2}}
	solution = 0
	CheckOrangesRotting(t, input, solution)

	input = [][]int{{0}}
	solution = 0
	CheckOrangesRotting(t, input, solution)

	input = [][]int{{1, 2}}
	solution = 1
	CheckOrangesRotting(t, input, solution)

	input = [][]int{
		{2, 1, 1},
		{1, 1, 1},
		{0, 1, 2}}
	solution = 2
	CheckOrangesRotting(t, input, solution)

}
