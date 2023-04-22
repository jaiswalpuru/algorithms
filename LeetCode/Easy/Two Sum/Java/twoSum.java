class Solution {
    public int[] twoSum(int[] nums, int target) {
        HashMap<Integer, Integer> numInd = new HashMap<>();
        for (int i=0; i<nums.length; i++) {
            int rem = target-nums[i];
            if (numInd.containsKey(rem)) return new int[]{i, numInd.get(rem)};
            numInd.put(nums[i], i);
        }
        return null;
    }
}
