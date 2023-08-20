class Solution {
    int[][] dir = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}};
    boolean[][] visited;
    int n, m;
    public int[][] floodFill(int[][] image, int sr, int sc, int color) {
        int c = image[sr][sc];
        n = image.length;
        m = image[0].length;
        visited = new boolean[image.length][image[0].length];
        dfs(image, sr, sc, c, color);
        return image;
    }

    private void dfs(int[][] image, int i, int j, int c, int color) {
        visited[i][j] = true;
        image[i][j] = color;
        for (int k=0; k<4; k++) {
            int x = i+dir[k][0];
            int y = j+dir[k][1];
            if (isSafe(x, y) && !visited[x][y] && image[x][y] == c)
                dfs(image, x, y, c, color);
        }
    }

    private boolean isSafe(int i, int j) {
        return i>=0 && j>=0 && i<n && j<m;
    }
}
