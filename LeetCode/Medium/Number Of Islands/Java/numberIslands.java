class Solution {
    int[][] dir = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}};
    int m, n;

    public int numIslands(char[][] grid) {
        m = grid.length;
        n = grid[0].length;
        int islands = 0;
        boolean[][] visited = new boolean[m][n];
        for (int i=0; i<m; i++) {
            for (int j=0; j<n; j++) {
                if (!visited[i][j] && grid[i][j] == '1') {
                    dfs(visited, grid, i, j);
                    islands++;
                }
            }
        }

        return islands;
    }

    private void dfs(boolean[][] visited, char[][] grid, int i, int j) {
        if (i<0 || j<0 || i>=m || j>=n || visited[i][j] || grid[i][j]=='0') return;

        visited[i][j]=true;
        for (int k=0; k<4; k++) {
            int x = dir[k][0] + i;
            int y = dir[k][1] + j;
            dfs(visited, grid, x, y);
        }
    }
}
