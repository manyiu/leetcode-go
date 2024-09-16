package reversebits

func reverseBits(num uint32) uint32 {
	var result uint32

	for i := 0; i < 32; i++ {
		result = result << 1

		if num&1 == 1 {
			result = result | 1
		}

		num = num >> 1
	}

	return result
}
