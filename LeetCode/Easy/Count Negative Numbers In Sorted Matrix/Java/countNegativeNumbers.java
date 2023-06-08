class Solution {
    public int countNegatives(int[][] grid) {
        int n = grid.length;
        int m = grid[0].length;
        int cnt = 0;
        for (int i=0; i<n; i++) {
            if (grid[i][m-1] < 0) {
                int j = m-1;
                while (j >= 0 && grid[i][j--] < 0) cnt++;
                
            }
        }
        return cnt;
    }
}
