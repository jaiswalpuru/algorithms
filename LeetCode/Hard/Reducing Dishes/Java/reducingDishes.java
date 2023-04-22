class Solution {
    public int maxSatisfaction(int[] satisfaction) {
        Arrays.sort(satisfaction);
        int[][] memo = new int[501][501];
        for (int i=0; i<=500; i++) Arrays.fill(memo[i], -1);
        return recur(satisfaction, 0, 1, memo);
    }

    int recur(int[] satis, int index, int time, int[][] memo) {
        if (index == satis.length) return 0;
        if (memo[index][time] != -1) return memo[index][time];
        memo[index][time] =  Math.max(satis[index]*time + recur(satis, index+1, time+1, memo), recur(satis, index+1, time, memo));
        return memo[index][time];
    }
}
