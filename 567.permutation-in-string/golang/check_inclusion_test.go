package check_inclusion

import "testing"

func CheckCheckInclusion(t *testing.T, input_s1 string, input_s2 string, solution bool) {
	check_inclusion := CheckInclusion(input_s1, input_s2)

	if check_inclusion == solution {
		t.Logf("OK")
	} else {
		t.Errorf("CheckInclusion(%s, %s): Got %t; expected %t", input_s1, input_s2, check_inclusion, solution)
		t.Fail()
	}
}

func TestCheckInclusion(t *testing.T) {
	input_s1 := "ab"
	input_s2 := "eidbaooo"
	CheckCheckInclusion(t, input_s1, input_s2, true)

	input_s2 = "eidboaoo"
	CheckCheckInclusion(t, input_s1, input_s2, false)

	input_s2 = "abadojfno"
	CheckCheckInclusion(t, input_s1, input_s2, true)
}
