class Solution {
    int[][] dir = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};

    public int findMaxFish(int[][] grid) {
        int maxFish = 0;
        int n = grid.length;
        int m = grid[0].length;
        boolean[][] visited = new boolean[n][m];
        
        for (int i=0; i<n; i++) {
            for (int j=0; j<m; j++) {
                if (grid[i][j] != 0 && !visited[i][j])
                    maxFish = Math.max(dfs(grid, i, j, visited), maxFish);
            }
        }

        return maxFish;
    }

    private int dfs(int[][] grid, int i, int j, boolean[][] visited) {        
        int fish = grid[i][j];
        visited[i][j] = true;    
        for (int k=0; k<4; k++) {
            int x = i + dir[k][0];
            int y = j + dir[k][1];
            if (x >=0 && y >= 0 && x < grid.length && y < grid[0].length && !visited[x][y] && grid[x][y] != 0) {
                fish += dfs(grid, x, y, visited);
            }
        }
        return fish;
    }
}
