class Solution {
    int[] parent;
    int[] size;

    public int makeConnected(int n, int[][] connections) {
        if (n-1 > connections.length) return -1;

        parent = new int[n];
        size = new int[n];
        for (int i=0; i<n; i++) {
            parent[i] = i;
            size[i] = 1;
        }

        for (int[] edge : connections) union(edge[0], edge[1]);

        for (int i=1; i<n; i++) {
            if (find(0) != find(i)){}
        }

        Set<Integer> s = new HashSet<>();
        for (int v : parent) s.add(v);
        return s.size()-1;
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
        if (x != parent[x]) {
            parent[x] = find(parent[x]);
        }
        return parent[x];
    }
}
