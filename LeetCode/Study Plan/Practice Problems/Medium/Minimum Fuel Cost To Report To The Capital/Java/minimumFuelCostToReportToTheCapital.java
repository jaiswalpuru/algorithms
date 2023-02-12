class Solution {
    public long minimumFuelCost(int[][] roads, int seats) {
        Map<Integer, List<Integer>> graph = new HashMap();
        int size = roads.length+1;
        int[] degree = new int[size];
        int[] subtree = new int[size];
        long fuel = 0;

        Arrays.fill(subtree, 1);

        for (int[] road : roads) {
            graph.computeIfAbsent(road[0], k -> new ArrayList<>());
            graph.computeIfAbsent(road[1], k -> new ArrayList<>());

            graph.get(road[0]).add(road[1]);
            graph.get(road[1]).add(road[0]);

            degree[road[0]]++;
            degree[road[1]]++;
        }

        ArrayDeque<Integer> queue = new ArrayDeque<>();

        for (int i=1; i<size; i++) {
            if (degree[i] == 1) queue.offer(i);
        }

        while(!queue.isEmpty()) {
            int curr = queue.poll();
            fuel += Math.ceil((double)subtree[curr]/seats);
            for (int neigh : graph.get(curr)) {
                degree[neigh]--;
                subtree[neigh] += subtree[curr];
                if (degree[neigh] == 1 && neigh != 0) queue.offer(neigh);
            }
        }

        return fuel;
    }
}
