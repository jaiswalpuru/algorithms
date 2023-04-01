class Solution {
    int n, m;
    int[][] dir = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}};
    public boolean exist(char[][] board, String word) {
        n = board.length;
        m = board[0].length;
        
        if (n == 1 && m == 1 && word.length() == 1) return word.charAt(0) == board[0][0];

        boolean[][] visited = new boolean[n][m];
        for (int i=0; i<n; i++) {
            for (int j=0; j<m; j++) {
                if (board[i][j] == word.charAt(0)) {
                    if (wordSearch(i, j, board, word, visited)) return true;
                }
            }
        }
        return false;
    }

    private boolean wordSearch(int i, int j, char[][] board, String word, boolean[][] visited) {
        if (word.length() == 0) return true;
        
        if (word.charAt(0) != board[i][j])
            return false;
    
        if (visited[i][j]) return false;

        visited[i][j] = true;
        boolean res = false;
        for (int k=0; k<4; k++) {
            int x = i+dir[k][0];
            int y = j+dir[k][1];
            if (isSafe(x, y)) {
                res = res || wordSearch(x, y,board, word.substring(1), visited);
            }
        }
        visited[i][j] = false;
        return res;
    }

    private boolean isSafe(int i, int j) {
        return i>=0 && j>=0 && i<n && j<m;
    }
}
