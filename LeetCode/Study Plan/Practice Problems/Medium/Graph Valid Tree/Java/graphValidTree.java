class Solution {
    int[] parent;
    int[] size;

    public boolean validTree(int n, int[][] edges) {
        parent = new int[n];
        size = new int[n];
        for (int i=0; i<n; i++) {
            parent[i] = i;
            size[i] = 1;
        }

        for (int[] e : edges) {
            if (!union(e[0], e[1])) return false;
        }
        for (int i=1; i<n; i++) {
            if (find(0) != find(i)) return false;
        }
        return true;
    }

    private int find(int x) {
        if (x != parent[x]) parent[x] = find(parent[x]);
        return parent[x];
    }

    private boolean union(int x, int y) {
        x = find(x);
        y = find(y);
        if (x == y) return false;
        if (size[x] > size[y]) {
            parent[y] = x;
            size[x] += size[y];
        }else {
            parent[x] = y;
            size[y] += size[x];
        }
        return true;
    }
}
