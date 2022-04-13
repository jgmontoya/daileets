package max_profit

import "testing"

func CheckMaxprofit(t *testing.T, input []int, solution int) {
	max_profit := maxProfit(input)
	if max_profit == solution {
		t.Log("OK")
	} else {
		t.Errorf("maxProfit(%d): Got %d; expected %d", input, max_profit, solution)
		t.Fail()
	}
}

func TestMaxProfit(t *testing.T) {
	input := []int{7, 1, 5, 3, 6, 4}
	solution := 5
	CheckMaxprofit(t, input, solution)

	input = []int{7, 6, 4, 3, 1}
	solution = 0
	CheckMaxprofit(t, input, solution)

	input = []int{4, 14, 1, 8, 2, 10}
	solution = 10
	CheckMaxprofit(t, input, solution)

	input = []int{3, 3, 5, 0, 0, 3, 1, 4}
	solution = 4
	CheckMaxprofit(t, input, solution)

	input = []int{2, 4, 1}
	solution = 2
	CheckMaxprofit(t, input, solution)

	input = []int{4, 7, 2, 1, 11}
	solution = 10
	CheckMaxprofit(t, input, solution)
}
