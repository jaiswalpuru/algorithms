class Solution {
    public int findMinArrowShots(int[][] points) {
        int cnt = 1;
        Arrays.sort(points, (a, b) -> Integer.compare(a[0], b[0]));
        int track = points[0][1];
        for (int i=1; i<points.length; i++) {
            if (track < points[i][0]) {
                cnt++;
                track = points[i][1];
            } else {
                track = Math.min(track, points[i][1]);
            }
        }
        return cnt;
    }
}
