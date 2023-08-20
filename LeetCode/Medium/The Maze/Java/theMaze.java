class Solution {
    int n, m;
    int[][] dir = {{0, -1}, {1, 0}, {0, 1}, {-1, 0}};
    public boolean hasPath(int[][] maze, int[] start, int[] destination) {
        n = maze.length;
        m = maze[0].length;
        boolean[][] visited = new boolean[n][m];
        return dfs(maze, start, destination, visited);
    }

    private boolean dfs(int[][] maze, int[] start, int[] dst, boolean[][] visited) {
        if (visited[start[0]][start[1]]) return false;
        if (start[0] == dst[0] && start[1] == dst[1]) return true;
        visited[start[0]][start[1]] = true;
        for (int i=0; i<4; i++) {
            int r = start[0], c = start[1];
            while(r >=0 && c >=0 && r < n && c < m && maze[r][c]==0) {
                r += dir[i][0];
                c += dir[i][1];
            }
            if (dfs(maze, new int[]{r-dir[i][0], c-dir[i][1]}, dst, visited)) return true;
        }
        return false;
    }
}
