class Solution {
    public List<List<Integer>> permute(int[] nums) {
        List<List<Integer>> res = new LinkedList();

        ArrayList<Integer> numList = new ArrayList<>();
        for (int val : nums) numList.add(val);

        int size = nums.length;
        backtrack(numList, size, 0, res);
        return res;
    }

    private void backtrack(ArrayList<Integer> nums, int n, int ind, List<List<Integer>> res) {
        if (ind == n) {
            res.add(new ArrayList<>(nums));
            return;
        }
        for (int i=ind; i<n; i++) {
            Collections.swap(nums, ind, i);
            backtrack(nums, n, ind+1, res);
            Collections.swap(nums, ind, i);
        }
    }
}
