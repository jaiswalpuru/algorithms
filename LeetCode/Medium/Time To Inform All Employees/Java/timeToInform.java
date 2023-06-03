class Solution {
    public int numOfMinutes(int n, int headID, int[] manager, int[] informTime) {
        Map<Integer, List<Integer>> graph = new HashMap<>();
        ArrayDeque<Integer> q = new ArrayDeque<>();
        for (int i=0; i<manager.length; i++) {
            if (manager[i] != -1)
                graph.computeIfAbsent(manager[i], k-> new ArrayList<>()).add(i);
        }
        q.offer(headID);
        int time = 0;
        while (!q.isEmpty()) {
            Integer curr = q.poll();
            time = Math.max(time, informTime[curr]);
            List<Integer> nei = graph.get(curr);
            if (nei != null) {
                for (int i=0; i<nei.size(); i++) {
                    q.offer(nei.get(i));
                    informTime[nei.get(i)] += informTime[curr];
                }
            }
        }

        return time;
    }
}
