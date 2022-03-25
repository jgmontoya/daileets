package flood_fill

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

func SliceEqual2D(first_slice [][]int, second_slice [][]int) bool {
	first_length := len(first_slice)
	second_length := len(second_slice)

	if first_length != second_length {
		return false
	}

	for i := 0; i < first_length; i++ {
		if !SliceEqual(first_slice[i], second_slice[i]) {
			return false
		}
	}
	return true
}

func CheckFloodFill(t *testing.T, input_image [][]int, sr int, sc int, newColor int, solution [][]int) {
	flood_fill := FloodFill(input_image, sr, sc, newColor)
	if SliceEqual2D(flood_fill, solution) {
		t.Log("OK")
	} else {
		t.Errorf("Got %d; expected %d", flood_fill, solution)
		t.Fail()
	}
}

func TestFloodFill(t *testing.T) {
	input_image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	sr := 1
	sc := 1
	newColor := 2
	solution := [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}
	CheckFloodFill(t, input_image, sr, sc, newColor, solution)

	input_image = [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	newColor = 1
	solution = [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	CheckFloodFill(t, input_image, sr, sc, newColor, solution)

	input_image = [][]int{{0, 0, 0}, {0, 0, 0}}
	sr = 0
	sc = 0
	newColor = 2
	solution = [][]int{{2, 2, 2}, {2, 2, 2}}
	CheckFloodFill(t, input_image, sr, sc, newColor, solution)

}
