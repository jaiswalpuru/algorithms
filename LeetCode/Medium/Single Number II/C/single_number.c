int singleNumber(int* nums, int numsSize){
    long lone = 0;
    for (int i=0; i<32; i++) {
        int bit_sum = 0;
        for (int j=0; j<numsSize; j++) bit_sum += (nums[j] >> i) & 1;
        long loner_bit = bit_sum%3;
        lone |= (loner_bit<<i);
    }
    return lone;
}
