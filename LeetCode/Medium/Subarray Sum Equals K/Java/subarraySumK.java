class Solution {
    public int subarraySum(int[] nums, int k) {
        Map<Integer, Integer> mp = new HashMap<>();
        int currSum = 0;
        int cnt = 0;
        for (int num : nums) {
            currSum += num;
            if (currSum == k) cnt++;
            cnt += mp.getOrDefault(currSum-k, 0);
            mp.put(currSum, mp.getOrDefault(currSum, 0)+1);
        }
        return cnt;
    }
}
