class Solution {
    public boolean checkStraightLine(int[][] coordinates) {
        int n = coordinates.length;
        int[] c1 = coordinates[0];
        int[] c2 = coordinates[1];
        for (int i=2; i<n; i++) {
            if ((coordinates[i][1]-c2[1]) * (c2[0]-c1[0]) != 
            (coordinates[i][0]-c2[0]) * (c2[1]-c1[1])) return false;
        }
        return true;
    }
}
