class Solution {
    int[][] dir = new int[][]{{0, 1}, {0, -1}, {1, 0}, {-1, 0}};

    public int maxDistance(int[][] grid) {
        int[][] visited = new int[grid.length][grid.length];
        ArrayDeque<Pair<Integer, Integer>> queue = new ArrayDeque<>();
        int size = grid.length;
        int distance = -1;

        for (int i=0; i<size; i++) {
            for (int j=0; j<size; j++) {
                visited[i][j] = grid[i][j];
                if (grid[i][j] == 1) {
                    queue.offer(new Pair(i, j));
                }
            }
        }

        while(!queue.isEmpty()) {
            int qLen = queue.size();

            while(qLen-- > 0) {
                Pair<Integer, Integer> curr = queue.poll();
                for (int[] d : dir) {
                    int x = d[0] + curr.getKey();
                    int y = d[1] + curr.getValue();
                    if (x >= 0 && y >= 0 && x < size && y < size && visited[x][y] == 0) {
                        queue.offer(new Pair(x, y));
                        visited[x][y] = 1;
                    }
                }
            }
            distance++;
        }

        return distance == 0 ? -1 : distance;
    }
}
