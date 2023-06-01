class Pair {
    int x, y;
    public Pair(int x, int y) {
        this.x = x;
        this.y = y;
    }
}

class Solution {
    public int shortestPathBinaryMatrix(int[][] grid) {
        int n = grid.length;
        if (grid[0][0] == 1) return -1;
        int[][] dir = {{1, 0}, {1, 1}, {0, 1}, {-1, 0}, {0, -1}, {-1, -1}, {1, -1}, {-1, 1}};
        boolean[][] visited = new boolean[n][n];
        ArrayDeque<Pair> q = new ArrayDeque<>();
        q.offer(new Pair(0, 0));
        int dis = 0;
        while(!q.isEmpty()) {
            int size = q.size();
            for (int i=0; i<size; i++) {
                Pair curr = q.poll();
                if (curr.x == n-1 && curr.y == n-1) return dis+1;
                for (int k=0; k<dir.length; k++) {
                    int x = curr.x + dir[k][0];
                    int y = curr.y + dir[k][1];
                    if (x < 0 || y < 0 || x >= n || y >= n || grid[x][y] == 1 || visited[x][y]) continue;
                    q.offer(new Pair(x, y));
                    visited[x][y] = true;
                }
            }
            dis++;
        }
        return -1;
    }
}
