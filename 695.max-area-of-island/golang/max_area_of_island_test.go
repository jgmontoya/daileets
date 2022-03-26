package max_area_of_island

import "testing"

func CheckMaxAreaOfIsland(t *testing.T, input [][]int, solution int) {
	max_area_of_island := MaxAreaOfIsland(input)
	if max_area_of_island == solution {
		t.Log("OK")
	} else {
		t.Errorf("Got %d; expected %d", max_area_of_island, solution)
		t.Fail()
	}
}

func TestMaxAreaOfIsland(t *testing.T) {
	input := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	solution := 6
	CheckMaxAreaOfIsland(t, input, solution)

	input = [][]int{{0, 0, 0, 0, 0, 0, 0, 0}}
	solution = 0
	CheckMaxAreaOfIsland(t, input, solution)

	input = [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 1, 1},
	}
	solution = 4
	CheckMaxAreaOfIsland(t, input, solution)
}
