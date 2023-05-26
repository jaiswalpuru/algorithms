class Solution {
    public int findCircleNum(int[][] isConnected) {
        int n = isConnected.length;
        boolean[] visited = new boolean[n];
        int cnt = 0;
        for (int i=0; i<n; i++) {
            if (!visited[i]) {
                cnt++;
                dfs(isConnected, visited, i);
            }
        }
        return cnt;
    }
    
    private void dfs(int[][] isConnected, boolean[] visited, int k) {
        visited[k] = true;
        for (int i=0; i<isConnected.length; i++) {
            if (isConnected[k][i] == 1 && !visited[i]) {
                dfs(isConnected, visited, i);
            }
        }
    }
    
}
