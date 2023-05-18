class Solution {
    public int longestOnes(int[] nums, int k) {
        int i=0, j=0, n = nums.length;
        int cnt = 0;
        int l = 0;
        while (j < n) {
            if (nums[j] == 0) {
                if (cnt == k) {
                    l = Math.max(l, j-i);
                    while(cnt >= k) {
                        if (nums[i] == 0)cnt--;
                        i++;
                    }
                }
                cnt++;
            }
            j++;
        }
        l = Math.max(l, j-i);
        return l;
    }
}
