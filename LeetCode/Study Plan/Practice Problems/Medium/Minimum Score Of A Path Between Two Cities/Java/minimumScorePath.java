// method 1
class Solution {
    public int minScore(int n, int[][] roads) {
        HashMap<Integer, List<Pair<Integer, Integer>>> g = new HashMap<>();
        
        //make graph
        for (int i=0;i<roads.length; i++) {
            g.computeIfAbsent(roads[i][0], k->new ArrayList<>());
            g.get(roads[i][0]).add(new Pair<>(roads[i][1], roads[i][2]));
            g.computeIfAbsent(roads[i][1], k->new ArrayList<>());
            g.get(roads[i][1]).add(new Pair<>(roads[i][0], roads[i][2]));
        }
 
        int minDis = Integer.MAX_VALUE;
        boolean[] visited = new boolean[n+1];
        ArrayDeque<Integer> q= new ArrayDeque<>();
        q.offer(1);
        while(!q.isEmpty()) {
            int from = q.poll();
            for (Pair<Integer, Integer> edge : g.get(from)) {
                minDis = Math.min(minDis, edge.getValue());
                if (visited[edge.getKey()]) continue;
                visited[edge.getKey()] = true;
                q.offer(edge.getKey());
            }
        }

        return minDis;
    }
}

//method 2
class Solution {
    int[] parent;
    int[] size;
    public int minScore(int n, int[][] roads) {
        int minDis = Integer.MAX_VALUE;
        parent = new int[n+1];
        size = new int[n+1];
        for (int i=1; i<=n; i++) {
            parent[i] = i;
            size[i] = 1;
        }

        for (int i=0; i<roads.length; i++) union(roads[i][0], roads[i][1]);

        for (int i=0; i<roads.length; i++) {
            int findN = find(n);
            if (findN == find(roads[i][0]) && findN == find(roads[i][1]))
                minDis = Math.min(minDis, roads[i][2]);
        }

        return minDis;
    }

    public void union(int x, int y) {
        x = find(x);
        y = find(y);
        if (x == y) return;
        if (size[x] > size[y]) {
            size[x] += size[y];
            parent[y] = x;
        }else {
            size[y] += size[x];
            parent[x] = y;
        }
    }

    public int find(int x) {
        if (x != parent[x]) parent[x] = find(parent[x]);
        return parent[x];
    }
}
