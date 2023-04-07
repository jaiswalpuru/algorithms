class Solution {
    public List<List<Integer>> combinationSum(int[] candidates, int target) {
        List<List<Integer>> res = new ArrayList<>();
        backtrack(candidates, target, 0, 0, new ArrayList<>(), res);
        return res;
    }

    private void backtrack(int[] candidates, int target, int currSum, int ind, List<Integer> set, List<List<Integer>> res) {
        if (currSum >= target) {
            if (currSum == target) res.add(new ArrayList<>(set));
            return;
        }

        for (int i=ind; i<candidates.length; i++) {
            if (currSum + candidates[i] > target) continue;
            set.add(candidates[i]);
            backtrack(candidates, target, currSum+candidates[i], i, set, res);
            set.remove(set.size()-1);
        }
    }
}
