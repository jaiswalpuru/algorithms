class P {
    int x, y;
    public P(int x, int y) {
        this.x = x;
        this.y = y;
    }
}

class Solution {
    int[][] dir = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}};

    public int orangesRotting(int[][] grid) {
        Queue<P> q = new LinkedList<>();
        int n = grid.length, m = grid[0].length;
        boolean[][] visited = new boolean[n][m];
        for (int i=0; i<n; i++) {
            for (int j=0; j<m; j++) {
                if (grid[i][j] == 2) q.add(new P(i, j));
            }
        }

        int min = -1;
    
        while(!q.isEmpty()) {
            int size = q.size();
            for (int i=0; i<size; i++) {
                P curr = q.poll();
                for (int k=0; k<4; k++) {
                    int x = curr.x+dir[k][0];
                    int y = curr.y+dir[k][1];
                    if (x < 0 || y < 0 || x >= n || y >= m || visited[x][y] || grid[x][y] == 0 || grid[x][y] == 2) continue;
                    grid[x][y] = 2;
                    visited[x][y] = true;
                    q.offer(new P(x, y));
                }
            }
            min++;
        }

        for (int i=0; i<n; i++)
            for (int j=0; j<m; j++) if (grid[i][j] == 1) return -1;

        return min == -1 ? 0 : min;
    }
}
