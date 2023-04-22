class Solution {
    public int canCompleteCircuit(int[] gas, int[] cost) {
        int size = gas.length;
        int currTank = 0, totalTank = 0;
        int start = 0;
        for (int i = 0; i < size; i++) {
            totalTank += gas[i]-cost[i];
            currTank += gas[i]-cost[i];
            if (currTank < 0) {
                start = i+1;
                currTank = 0;
            }
        }
        if (totalTank < 0) return -1;
        return start;
    }
}
