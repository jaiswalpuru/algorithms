class Solution {
    HashMap<Integer, List<Integer>> g;
    boolean[] visited;
    int cnt = 0;
    public int minReorder(int n, int[][] connections) {
        g = new HashMap<>();
        visited = new boolean[n];

        for (int[] e : connections) {
            int u = e[0];
            int v = e[1];
            g.computeIfAbsent(u, k->new ArrayList<>());
            g.get(u).add(-v);
            g.computeIfAbsent(v, k->new ArrayList<>());
            g.get(v).add(u);
        }

        for (int i=0; i<n; i++) {
            if (!visited[i]) {
                dfs(i);
            }
        }

        return cnt;
    }

    public void dfs(int u) {
        visited[u] = true;
        List<Integer> neighbors = g.get(u);
        if (neighbors != null)
            for (Integer nei : neighbors) {
                if (!visited[Math.abs(nei)]) {
                    if (nei < 0) cnt++;
                    dfs(Math.abs(nei));
                }
            }
    }
}
