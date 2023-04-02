class Solution {
    public int longestConsecutive(int[] nums) {
        HashSet<Integer> set = new HashSet<>();
        int n = nums.length;
        int res = 0;

        for (int i=0; i<n; i++) set.add(nums[i]);

        for (int i=0; i<n; i++) {
            if (!set.contains(nums[i]-1)) {
                int val = nums[i];
                int cnt = 0;
                while(set.contains(val)) {
                    val++;
                    cnt++;
                }
                res = Math.max(res, cnt);
            }
        }

        return res;
    }
}
