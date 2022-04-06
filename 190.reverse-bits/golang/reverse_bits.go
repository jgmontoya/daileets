package reverse_bits

func reverseBits(num uint32) uint32 {
	var reversed_bits uint32 = 0
	var mask_checker uint32 = 1
	var mask_applier uint32 = 1 << 31
	for i := 0; i < 32; i++ {
		if num&mask_checker > 0 {
			reversed_bits = reversed_bits | mask_applier
		}
		mask_checker = mask_checker << 1
		mask_applier = mask_applier >> 1
	}
	return reversed_bits
}
