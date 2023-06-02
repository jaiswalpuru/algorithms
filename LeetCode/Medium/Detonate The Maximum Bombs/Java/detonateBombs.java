class Solution {
    Map<Integer, List<Integer>> graph;
    public int maximumDetonation(int[][] bombs) {
        graph = new HashMap<>();
        int n = bombs.length;
        for (int i=0; i<n; i++) {
            for (int j=0; j<n; j++) {
                if (i != j) {
                    if (isEuclidDistanceGood(bombs[i], bombs[j])) {
                        graph.computeIfAbsent(i, k -> new ArrayList<>()).add(j);
                    }
                }
            }
        }
        int res = 0;
        for (int i=0; i<n; i++) {
            boolean[] visited = new boolean[n];
            res = Math.max(res, dfs(i, visited));
        }
        return res;
    }

    private int dfs(int src, boolean[] visited) {
        visited[src] = true;
        int res = 1;
        List<Integer> nei = graph.get(src);
        if (nei != null)
            for (Integer v : nei) {
                if (!visited[v]) res += dfs(v, visited);
            }
        return res;
    }

    private boolean isEuclidDistanceGood(int[] b1, int[] b2) {
        int x1 = b1[0];
        int x2 = b2[0];
        int y1 = b1[1];
        int y2 = b2[1];
        int radius = b1[2];
        return Math.sqrt(Math.pow(x2-x1, 2)+Math.pow(y2-y1, 2)) <= radius;
    }
}
