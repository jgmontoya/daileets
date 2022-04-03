package climb_stairs

import "testing"

func CheckClimbStairs(t *testing.T, input int, solution int) {
	climb_stairs := climbStairs(input)
	if climb_stairs == solution {
		t.Log("OK")
	} else {
		t.Errorf("Got %d; expected %d", climb_stairs, solution)
		t.Fail()
	}
}
func TestClimbStairs(t *testing.T) {
	input := 1
	solution := 1
	CheckClimbStairs(t, input, solution)

	input = 2
	solution = 2
	CheckClimbStairs(t, input, solution)

	input = 3
	solution = 3
	CheckClimbStairs(t, input, solution)

	input = 4
	solution = 5
	CheckClimbStairs(t, input, solution)
}
