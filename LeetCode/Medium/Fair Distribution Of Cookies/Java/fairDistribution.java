class Solution {
    int res;
    public int distributeCookies(int[] cookies, int k) {
        res = Integer.MAX_VALUE;
        int[] cook = new int[k];
        backtrack(cookies, k, 0, cook);
        return res;
    }

    private void backtrack(int[] cookies, int k, int ind, int[] cook) {
        if (ind == cookies.length) {
            int temp = Integer.MIN_VALUE;
            for (int i=0; i<k; i++) temp = Math.max(temp, cook[i]);
            res = Math.min(res, temp);
            return;
        }

        for (int i=0; i<k; i++) {
            cook[i] += cookies[ind];
            backtrack(cookies, k, ind+1, cook);
            cook[i] -= cookies[ind];
        }
    }
}
