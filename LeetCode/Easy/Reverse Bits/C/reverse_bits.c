uint32_t reverseBits(uint32_t n) {
    uint32_t ret = 0;
    uint32_t pow = 31;
    while(n != 0) {
        ret += (n&1)<<pow;
        n >>= 1;
        pow -= 1;
    }
    return ret;
}
