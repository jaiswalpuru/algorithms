/**
 * Definition for a category handler.
 * class CategoryHandler {
 *     public CategoryHandler(int[] categories);
 *     public boolean haveSameCategory(int a, int b);
 * };
 */
class Solution {
    
	public int numberOfCategories(int n, CategoryHandler categoryHandler) {
    	UnionFind uf = new UnionFind(n);
        for (int i=0; i<n; i++) {
            for (int j=i+1; j<n; j++) {
                if (categoryHandler.haveSameCategory(i, j)) uf.union(i, j);
            }
        }
        return uf.getCount();
	}

    private class UnionFind {
        int[] parent;
        int[] size;
        int cnt; // to count the number of components
        UnionFind(int n) {
            parent = new int[n];
            size = new int[n];
            for (int i=0; i<n; i++) {
                parent[i] = i;
                size[i] = 1;
            }
            cnt = n;
        }

        public int find(int x) {
            if (x != parent[x]) parent[x] = find(parent[x]);
            return parent[x];
        }

        public void union(int x, int y) {
            x = find(x);
            y = find(y);
            if (x == y) return;
            cnt--;
            if (size[x] > size[y]) {
                parent[y] = x;
                size[x] += size[y];
            } else {
                parent[x] = y;
                size[y] += size[x];
            }
        }

        public int getCount() { return cnt; }
    }
}


// ----------------- using dfs -------------------
/**
 * Definition for a category handler.
 * class CategoryHandler {
 *     public CategoryHandler(int[] categories);
 *     public boolean haveSameCategory(int a, int b);
 * };
 */
class Solution {

    private void dfs(List<Integer>[] graph, boolean[] visited, int src) {
        visited[src] = true;
        for (int i=0; i<graph[src].size(); i++) {
            if (!visited[graph[src].get(i)]) dfs(graph, visited, graph[src].get(i));
        }
    }

	public int numberOfCategories(int n, CategoryHandler categoryHandler) {
    	List<Integer>[] graph = new ArrayList[n];
        for (int i=0; i<n; i++) graph[i] = new ArrayList<>();
        for (int i=0; i<n; i++) {
            for (int j=i+1; j<n; j++) {
                if (categoryHandler.haveSameCategory(i, j)) {
                    graph[i].add(j);
                    graph[j].add(i);
                }
            }
        }

        int components = 0;
        boolean[] visited = new boolean[n];
        for (int i=0; i<n; i++) {
            if (!visited[i]) {
                components++;
                dfs(graph, visited, i);
            }
        }
        return components;
	}
}
