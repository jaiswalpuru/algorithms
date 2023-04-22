class Solution {
    int[] parent;
    int[] size;

    public void union(int x, int y) {
        x = find(x);
        y = find(y);
        if (x == y) return;
        if (size[x] > size[y]) {
            size[x] += size[y];
            parent[y] = x;
        } else {
            size[y] += size[x];
            parent[x] = y;
        }
    }

    public int find(int x) {
        if (x != parent[x]) parent[x] = find(parent[x]);
        return parent[x];
    }

    public long countPairs(int n, int[][] edges) {
        parent = new int[n];
        size = new int[n];
        for (int i=0; i<n; i++) {
            parent[i] = i;
            size[i] = 1;
        }

        for (int[] e : edges) union(e[0], e[1]);

        HashMap<Integer, Long> count = new HashMap<>();
        for (int i=0; i<n; i++) count.put(find(i), count.getOrDefault(find(i), 0L)+1);

        long totalPairs = (long)n*(n-1)/2;
        System.out.println(totalPairs);
        for (Integer key : count.keySet()) {
            long val = count.get(key);
            totalPairs -= val*(val-1)/2;
        }
        return totalPairs;
    }
}
