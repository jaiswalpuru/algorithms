class Solution {
    public int minJumps(int[] arr) {
        int minJump = 0;
        int size = arr.length;
        Map<Integer, List<Integer>> graph = new HashMap<>();
        Set<Integer> visited = new HashSet<>();
        List<Integer> q = new LinkedList<>();
        q.add(0);

        for (int i = 0; i < size; i++) {
            graph.computeIfAbsent(arr[i], v -> new LinkedList<>()).add(i);
        }

        while (!q.isEmpty()) {
            List<Integer> nex = new LinkedList<>();
            for (int curr : q) {
                if (curr == size - 1) {
                    return minJump;
                }

                for (int child : graph.get(arr[curr])) {
                    if (!visited.contains(child)) {
                        visited.add(child);
                        nex.add(child);
                    }
                }

                graph.get(arr[curr]).clear();

                if (curr + 1 < size && !visited.contains(curr + 1)) {
                    visited.add(curr + 1);
                    nex.add(curr + 1);
                }
                if (curr - 1 >= 0 && !visited.contains(curr - 1)) {
                    visited.add(curr - 1);
                    nex.add(curr - 1);
                }
            }
            q = nex;
            minJump++;
        }

        return -1;
    }
}
