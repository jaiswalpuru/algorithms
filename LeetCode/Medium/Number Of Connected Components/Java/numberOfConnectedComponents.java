class Solution {
    int[] parent;
    int[] size;

    public int countComponents(int n, int[][] edges) {
        parent = new int[n];
        size = new int[n];
        
        for (int i=0; i<n; i++) {
            parent[i] = i;
            size[i] = 1;
        }

        for (int[] e : edges) union(e[0], e[1]);

        HashSet<Integer> set = new HashSet<>();
        for (int i=0; i<n; i++) {
            int p = find(i);
            set.add(p);
        }
        return set.size();
    }

    private int find(int x) {
        if (x != parent[x]) parent[x] = find(parent[x]);
        return parent[x];
    }

    private void union(int x, int y) {
        x = find(x);
        y = find(y);
        if (x == y) return;
        if (size[x] > size[y]) {
            parent[y] = x;
            size[x] += size [y];
        }else {
            parent[x] = y;
            size[y] += size[x];
        }
    }
}
