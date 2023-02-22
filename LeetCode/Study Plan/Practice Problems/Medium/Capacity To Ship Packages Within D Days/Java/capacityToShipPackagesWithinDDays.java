class Solution {
    public int shipWithinDays(int[] weights, int days) {
        int totalWt = 0;
        int maxWt = weights[0];

        for (int wt : weights) {
            maxWt = Math.max(maxWt, wt);
            totalWt += wt;
        }

        int low = maxWt;
        int high = totalWt;

        while (low < high) {
            int mid = (low+high)/2;
            if (isPossible(mid, days, weights)) {
                high = mid;
            } else {
                low = mid+1;
            }
        }

        return low;
    }

    public boolean isPossible(int capacity, int days, int[] weights) {
        int totalDays = 1;
        int totalWt = 0;
        int size =weights.length;
        
        for (int i=0; i<size; i++) {
            totalWt += weights[i];
            if (totalWt > capacity) {
                totalDays+=1;
                totalWt = weights[i];
            }
        }

        return totalDays <= days;
    }
}
