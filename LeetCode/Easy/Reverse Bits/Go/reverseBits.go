func reverseBits(num uint32) uint32 {
	ret := uint32(0)
	pow := uint32(31)
	for num != 0 {
		ret += (num & 1) << pow
		num >>= 1
		pow--
	}
	return ret
}
