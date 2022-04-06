package hamming_weight

func hammingWeight(num uint32) int {
	var hamming_weight uint32 = 0
	var mask uint32 = 1
	for i := 0; i < 32; i++ {
		if num&mask > 0 {
			hamming_weight++
		}
		mask = mask << 1
	}
	return int(hamming_weight)
}
