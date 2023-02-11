class Solution {
    // 0 -> represents red and 1 -> represents blue
    public int[] shortestAlternatingPaths(int n, int[][] redEdges, int[][] blueEdges) {
        Map<Integer, List<Pair<Integer, Integer>>> graph= new HashMap<>();
        int sizeRed = redEdges.length;
        int blueSize = blueEdges.length;

        // make the graph with red edges
        for (int i=0; i<sizeRed; i++) {
            graph.computeIfAbsent(redEdges[i][0], k -> new ArrayList<>()).add(new Pair(redEdges[i][1], 0));
        }

        // make the graph with blue edges
        for (int i=0; i<blueSize; i++) {
            graph.computeIfAbsent(blueEdges[i][0], k -> new ArrayList<>()).add(new Pair(blueEdges[i][1], 1));
        }

        Map<Pair<Integer, Integer>, Boolean> visited = new HashMap<>();
        ArrayDeque<Pair<Integer, Integer>> queue = new ArrayDeque<>();

        queue.offer(new Pair(0, -1));

        int[] res = new int[n];
        Arrays.fill(res, -1);
        res[0] = 0;
        int dis = 0;

        while (!queue.isEmpty()) {
            int qLen = queue.size();
            for (int i=0; i<qLen; i++) {
                Pair<Integer, Integer> curr = queue.poll();
                int node = curr.getKey();
                int color = curr.getValue();
                if (visited.containsKey(curr)) continue;
                if (res[node] == -1) {
                    res[node] = dis;
                }else {
                    res[node] = Math.min(res[node], dis);
                }
                visited.put(curr, true);
                if (graph.get(node) == null) {
                    continue;
                }
                for (Pair<Integer, Integer> adj : graph.get(node)) {
                    if (!visited.containsKey(adj)) {
                        int v = adj.getKey();
                        int c = adj.getValue();
                        if (c != color) {
                            queue.offer(new Pair(v, c));
                        }
                    }
                }
            }
            dis++;
        }
        return res;
    }
}
