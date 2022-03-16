package first_bad_version

import "testing"

func TestFirstBadVersion(t *testing.T) {
	t.Log("should return first true version")
	versions := 5
	isBadVersion = func(n int) bool {
		isBad := []bool{false, false, false, true, true}[n-1]
		return isBad
	}

	first_bad_version := FirstBadVersion(versions)
	if first_bad_version != 4 {
		t.Errorf("Got %d; expected 4", first_bad_version)
		t.Fail()
	} else {
		t.Log("OK")
	}
}
