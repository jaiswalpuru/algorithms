class Solution {
    public boolean containsDuplicate(int[] nums) {
        HashMap<Integer, Integer> count = new HashMap<>();
        for (int val : nums) {
            count.put(val, count.getOrDefault(val, 0)+1);
        }
        return count.size() != nums.length;
    }
}
