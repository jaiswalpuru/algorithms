class Solution {
    public int orangesRotting(int[][] grid) {
        int n = grid.length, m = grid[0].length;
        boolean[][] visited = new boolean[n][m];

        int minTime = -1;
        ArrayDeque<Map.Entry<Integer, Integer>> q = new ArrayDeque<>();
        for (int i=0; i<n; i++) {
            for (int j=0; j<m; j++) {
                if (grid[i][j] == 2) {
                    q.offer(Map.entry(i, j));
                }
            }
        }

        int[][] dir  = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}};
        while (!q.isEmpty()) {
            int size = q.size();
            while(size-- > 0) {
                Map.Entry<Integer, Integer> mp = q.poll();
                int i = mp.getKey(), j = mp.getValue();
                for (int k=0; k<4; k++) {
                    int r = i + dir[k][0];
                    int c = j + dir[k][1];
                    if (r < 0 || c < 0 || r >= n || c >= m || visited[r][c] || grid[r][c] == 0 || grid[r][c] == 2) continue;
                    grid[r][c] = 2;
                    visited[r][c] = true;
                    q.offer(Map.entry(r, c));
                }
            }
            minTime++;
        }

        for (int i=0; i<n; i++) {
            for (int j=0; j<m; j++) {
                if (grid[i][j] == 1) return -1;
            }
        }
        return minTime == -1 ? 0 : minTime;
    }
}
