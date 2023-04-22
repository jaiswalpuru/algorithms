class Solution {
    int n;
    int[] days, costs;
    int[] daysPass = {1, 7, 30};
    int[] memo;
    public int mincostTickets(int[] days, int[] costs) {
        this.n = days.length;
        this.days = days;
        this.costs = costs;
        this.memo = new int[n];
        Arrays.fill(this.memo, -1);
        return recur(0);
    }

    public int recur(int curInd) {
        if (curInd >= n) return 0;
        
        if (memo[curInd] != -1) return memo[curInd];

        int minCost = (int)1e7;
        for (int i=0; i<3; i++) {
            int nextInd = getNextIndex(curInd, daysPass[i]);
            
            int currCost = costs[i] + recur(nextInd);
            minCost = Math.min(minCost, currCost);
        }
        memo[curInd] = minCost;
        return minCost;
    }

    public int getNextIndex(int curInd, int day) {
        int d = days[curInd]+day;
        for (int i=curInd; i<n; i++) {
            if (days[i] >= d) {
                return i;
            }
        }
        return n;
    }
}
