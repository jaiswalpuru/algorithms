class Solution {
    int m, n;
    int[][] dir = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}};
    public int closedIsland(int[][] grid) {
        m = grid.length;
        n = grid[0].length;
        boolean[][] visited = new boolean[m][n];
        int cnt=0;
        for (int i=1; i<m-1; i++) {
            for (int j=1; j<n-1; j++) {
                if (grid[i][j] == 0 && !visited[i][j]) {
                    if (dfs(grid, i, j, visited)){
                        cnt++;
                    }
                }
            }
        }
        return cnt;
    }

    private boolean dfs(int[][] grid, int i, int j, boolean[][] visited) {
        if (i==0 || j==0 || i==m-1 || j==n-1 || visited[i][j]) return false;

        visited[i][j] = true;
        boolean ans = true;
        for (int k=0; k<4; k++) {
            int x = i+dir[k][0];
            int y = j+dir[k][1];
            if (x >=0 && y>=0 && grid[x][y]==0 && !visited[x][y])
                if (!dfs(grid, x, y, visited)) ans = false;
        }
        return ans;
    }
}
