package hamming_weight

import "testing"

func CheckHammingWeight(t *testing.T, input uint32, solution int) {
	hamming_weight := hammingWeight(input)
	if hamming_weight == solution {
		t.Logf("OK")
	} else {
		t.Logf("hammingWeight(%d); Got %d, expected %d", input, hamming_weight, solution)
		t.Fail()
	}
}
func TestHammingWeight(t *testing.T) {
	var input uint32 = 11
	solution := 3
	CheckHammingWeight(t, input, solution)

	input = 128
	solution = 1
	CheckHammingWeight(t, input, solution)

	input = 4294967293
	solution = 31
	CheckHammingWeight(t, input, solution)
}
