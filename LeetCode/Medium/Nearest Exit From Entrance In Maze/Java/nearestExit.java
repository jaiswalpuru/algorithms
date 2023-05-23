class P {
    int x;
    int y;
    int dis;
    public P(int x, int y, int d) {
        this.x = x;
        this.y = y;
        this.dis = d;
    }
}

class Solution {
    int[][] dir = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}};

    public int nearestExit(char[][] maze, int[] entrance) {
        int n = maze.length;
        int m = maze[0].length;
        boolean[][] visited = new boolean[n][m];
        Queue<P> q = new LinkedList<>();
    
        for (int k=0; k<4; k++) {
            int x = entrance[0] + dir[k][0];
            int y = entrance[1] + dir[k][1];
            if (x < 0 || y < 0 || x >= n || y >= m || maze[x][y] == '+') continue;
            q.offer(new P(x, y, 1));
        }
        visited[entrance[0]][entrance[1]] = true;
        while(!q.isEmpty()) {
            int size = q.size();
            for (int i=0; i<size; i++) {
                P curr = q.poll();
                if (curr.x == 0 || curr.y == 0 || curr.x == n-1 || curr.y == m-1) return curr.dis;
                if (visited[curr.x][curr.y]) continue;
                visited[curr.x][curr.y] = true;
                for (int k=0; k<4; k++) {
                    int x = curr.x + dir[k][0];
                    int y = curr.y + dir[k][1];
                    if (x < 0 || y < 0 || x >= n || y >= m || maze[x][y] == '+' || visited[x][y]) continue;
                    q.offer(new P(x, y, curr.dis+1));
                }
            }
        }

        return -1;
    }
}
