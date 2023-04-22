class Solution {
    public List<List<Integer>> combinationSum2(int[] candidates, int target) {
        List<List<Integer>> res = new ArrayList<>();
        Arrays.sort(candidates);
        backtrack(candidates, target, res, 0, 0, new ArrayList<>());
        return res;
    }

    private void backtrack(int[] cand, int target, List<List<Integer>> res, int ind, int sum, List<Integer> temp) {
        if (sum == target) {
            res.add(new ArrayList<>(temp));
            return;
        }

        for (int i=ind; i<cand.length; i++) {
            if (i > ind && cand[i] == cand[i-1]) continue;
            if (sum + cand[i] > target) continue;
            temp.add(cand[i]);
            backtrack(cand, target, res, i+1, sum+cand[i], temp);
            temp.remove(temp.size()-1);
        }
    }
}
