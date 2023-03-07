class Solution {
    public long minimumTime(int[] time, int totalTrips) {
        Arrays.sort(time);

        int size = time.length;
        long l = 1;
        long r = (long)time[size-1]*totalTrips;
        long res = 0;

        while(l <= r) {
            long mid = l+(r-l)/2;
            if (isPossible(time, mid, totalTrips)) {
                res = mid;
                r = mid-1;
            }else {
                l = mid+1;
            }
        }

        return res;
    }

    public boolean isPossible(int[] time, long t, int totalTrips) {
        long trip = 0;

        for (int i = 0; i<time.length; i++) {
            trip += t/(long)time[i];
        }

        return trip >= totalTrips;
    }
}
