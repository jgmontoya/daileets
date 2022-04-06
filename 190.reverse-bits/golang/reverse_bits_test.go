package reverse_bits

import "testing"

func CheckReverseBits(t *testing.T, input uint32, solution uint32) {
	reversed_bits := reverseBits(input)
	if reversed_bits == solution {
		t.Logf("OK")
	} else {
		t.Logf("reverseBits(%d); Got %d, expected %d", input, reversed_bits, solution)
		t.Fail()
	}
}
func TestReverseBits(t *testing.T) {
	var input uint32 = 43261596
	var solution uint32 = 964176192
	CheckReverseBits(t, input, solution)

	input = 4294967293
	solution = 3221225471
	CheckReverseBits(t, input, solution)
}
