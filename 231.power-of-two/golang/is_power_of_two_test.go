package is_power_of_two

import "testing"

func CheckIsPowerOfTwo(t *testing.T, input int, solution bool) {
	is_power_of_two := isPowerOfTwo(input)
	if is_power_of_two == solution {
		t.Logf("OK")
	} else {
		t.Errorf("isPowerOfTwo(%d): Got %t; expected %t", input, is_power_of_two, solution)
		t.Fail()
	}
}

func TestIsPowerOfTwo(t *testing.T) {
	input := 0
	solution := false
	CheckIsPowerOfTwo(t, input, solution)

	input = 1
	solution = true
	CheckIsPowerOfTwo(t, input, solution)

	input = 2
	solution = true
	CheckIsPowerOfTwo(t, input, solution)

	input = 3
	solution = false
	CheckIsPowerOfTwo(t, input, solution)

	input = -16
	solution = false
	CheckIsPowerOfTwo(t, input, solution)

	input = 16
	solution = true
	CheckIsPowerOfTwo(t, input, solution)

	input = 6
	solution = false
	CheckIsPowerOfTwo(t, input, solution)

}
