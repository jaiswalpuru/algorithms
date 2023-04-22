class Solution {
    public List<List<Integer>> subsets(int[] nums) {
        List<List<Integer>> res = new ArrayList<>();
        backtrack(nums, 0, res, new ArrayList<>());
        return res;
    }

    private void backtrack(int[] nums, int ind, List<List<Integer>> res, List<Integer> temp) {
        res.add(new ArrayList<>(temp));

        for (int i=ind; i<nums.length; i++) {
            temp.add(nums[i]);
            backtrack(nums, i+1, res, temp);
            temp.remove(temp.size()-1);
        }
    }
}
