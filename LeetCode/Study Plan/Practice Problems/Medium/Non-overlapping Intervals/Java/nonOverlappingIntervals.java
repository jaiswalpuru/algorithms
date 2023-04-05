class Solution {
    public int eraseOverlapIntervals(int[][] intervals) {
        int n = intervals.length;
        Arrays.sort(intervals, new Comparator<int[]>() {
            public int compare(int[] a, int[] b) {
                return a[0] - b[0];
            }
        });
        Arrays.sort(intervals, new Comparator<int[]>() {
            public int compare(int[] a, int[] b) {
                return a[1] - b[1];
            }
        });
    
        int remove = 0;
        int j=0;
        for (int i=1; i<n; i++) {
            if (intervals[i][0] < intervals[j][1]) {
                remove++;
            } else {
                j=i;
            }
        }
        return remove;
    }
}
