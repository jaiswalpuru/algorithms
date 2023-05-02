class Solution {
    public int arraySign(int[] nums) {
        int cnt = 0;
        for (int val : nums) {
            if (val == 0) return 0;
            if (val < 0) cnt++;
        }
        return cnt%2 == 0 ? 1 : -1;
    }
}
