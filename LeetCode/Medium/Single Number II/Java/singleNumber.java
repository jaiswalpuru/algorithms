class Solution {
    public int singleNumber(int[] nums) {
        int lone = 0;
        for (int s=0; s<32; s++) {
            int bitSum = 0;
            for (int num : nums) bitSum += (num >> s) & 1;
            int lonerBit = bitSum%3;
            lone |= (lonerBit << s);
        }
        return lone;
    }
}
