class Solution {
    List<List<Integer>> res;
    public List<List<Integer>> combine(int n, int k) {
        res = new ArrayList<>();
        backtrack(n, k, new ArrayList<>(), 1);
        return res;
    }

    private void backtrack(int n, int k, List<Integer> list, int start) {
        if (k == 0) {
            res.add(new ArrayList<>(list));
            return;
        }
        for (int i=start; i<=n; i++) {
            list.add(i);
            backtrack(n, k-1, list, i+1);
            list.remove(list.size()-1);
        }
    }
}
