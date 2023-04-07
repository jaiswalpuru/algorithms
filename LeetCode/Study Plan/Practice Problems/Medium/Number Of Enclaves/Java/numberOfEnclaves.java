class Solution {
    int m;
    int n;
    int[][] dir = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}};
    int t;
    public int numEnclaves(int[][] grid) {
        m = grid.length;
        n = grid[0].length;
        boolean[][] visited = new boolean[m][n];
        int enclaves = 0;

        for (int i=1; i<m-1; i++) {
            for (int j=1; j<n-1; j++) {
                if (grid[i][j] == 1 && !visited[i][j]) {
                    t = 0;
                    if (dfs(i, j, grid, visited)) {
                        enclaves += t;
                    }
                }
            }
        }
        return enclaves;
    }

    private boolean dfs(int i, int j, int[][] grid, boolean[][] visited) {
        if (i==0 || j==0 || i==m-1 || j==n-1) return false;
        visited[i][j] = true;
        boolean ans = true;
        t++;
        for (int k=0; k<4; k++) {
            int x = i+dir[k][0];
            int y = j+dir[k][1];
            if (x>=0 && y>=0 && x<m && y<n && grid[x][y] == 1 && !visited[x][y]) {
                if (!dfs(x, y, grid, visited)) ans = false;
            }
        }
        return ans;
    }
}
