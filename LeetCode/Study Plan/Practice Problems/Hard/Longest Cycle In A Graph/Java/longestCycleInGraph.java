class Solution {
    int res = -1; 

    public void dfs(int u, int[] edges, Map<Integer, Integer> distance, boolean[] visited) {
        visited[u] = true;
        int nei = edges[u];
        if (nei != -1 && !visited[nei]) {
            distance.put(nei, distance.get(u)+1);
            dfs(nei, edges, distance, visited);
        }else if (nei != -1 && distance.containsKey(nei)){
            res = Math.max(res, distance.get(u)-distance.get(nei)+1);
        }
    }

    public int longestCycle(int[] edges) {
        int n = edges.length;
        boolean[] visited = new boolean[n];

        for (int i=0; i<n; i++) {
            if (!visited[i]) {
                Map<Integer, Integer> distance = new HashMap<>();
                distance.put(i,1);
                dfs(i, edges, distance, visited);
            }
        }

        return res;
    }

}
