class Solution {
    public List<List<Integer>> combinationSum3(int k, int n) {
        List<List<Integer>> res = new ArrayList<>();
        backtrack(res, new ArrayList<>(), k, n, 1, 0);
        return res;
    }

    private void backtrack(List<List<Integer>> res, List<Integer> set, int k, int n, int start, int sum) {
        if (set.size() == k) {
            if (sum == n)
                res.add(new ArrayList<>(set));
            return;
        }

        for (int i=start; i<=9; i++) {
            if (i+sum <= n) {
                set.add(i);
                backtrack(res, set, k, n, i+1, sum+i);
                set.remove(set.size()-1);
            }
        }
    }
}
