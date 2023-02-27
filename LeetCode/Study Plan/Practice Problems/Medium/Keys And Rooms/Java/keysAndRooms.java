class Solution {
    public boolean canVisitAllRooms(List<List<Integer>> rooms) {
        int size = rooms.size();
        ArrayDeque<Integer> queue = new ArrayDeque<>();
        queue.offer(0);
        HashMap<Integer, Boolean> visited = new HashMap<>();

        while (!queue.isEmpty()) {
            int qSize = queue.size();
            
            for (int i=0; i<qSize; i++) {
                int key = queue.poll();
                visited.put(key, true);
                for (int j=0; j<rooms.get(key).size(); j++) {
                    if (!visited.getOrDefault(rooms.get(key).get(j), false)) {
                        queue.offer(rooms.get(key).get(j));
                    }
                }
            }
        }

        for (int i=0; i<size; i++) {
            if (!visited.getOrDefault(i, false)) {
                return false;
            }
        }

        return true;
    }
}
