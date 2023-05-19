enum Color {
    RED,
    BLUE
}

class Solution {
    public boolean isBipartite(int[][] graph) {
        int n = graph.length;
        HashMap<Integer, Color> color = new HashMap<>();
        for (int i=0; i<n; i++) {
            if (!color.containsKey(i)) {
                color.put(i, Color.BLUE);
                ArrayDeque<Integer> q = new ArrayDeque<>();
                q.add(i);
                while (!q.isEmpty()) {
                    int u = q.poll();
                    for (int j=0; j<graph[u].length; j++) {
                        int v = graph[u][j];
                        if (!color.containsKey(v)) {
                            if (color.get(u).compareTo(Color.BLUE)==0)
                                color.put(v, Color.RED);
                            else color.put(v, Color.BLUE);
                            q.add(v);
                        } else {
                            if (color.get(u) == color.get(v))
                                return false;
                        }
                    }
                }
            }
        }

        return true;
    }
}
