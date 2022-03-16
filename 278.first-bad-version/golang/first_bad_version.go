package first_bad_version

var isBadVersion = func(n int) bool {
	return true
}

func FirstBadVersion(n int) int {
	return RecursiveFirstBadVersion(n, 0)
}

func RecursiveFirstBadVersion(n int, accumulated_version int) int {
	if n == 1 {
		return n + accumulated_version
	}

	mid_version := (n / 2)
	if isBadVersion(mid_version + accumulated_version) {
		return RecursiveFirstBadVersion(mid_version, accumulated_version)
	} else {
		return RecursiveFirstBadVersion(mid_version+1, accumulated_version+mid_version)
	}
}
