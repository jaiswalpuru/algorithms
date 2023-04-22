class Solution {
    public boolean canReach(int[] arr, int start) {
        HashMap<Integer, Boolean> visited = new HashMap<>();
        int size = arr.length;
        ArrayDeque<Integer> queue = new ArrayDeque<>();
        
        queue.offer(start);
        while(!queue.isEmpty()) {
            int curr = queue.poll();
            if (arr[curr] == 0) return true;

            if (curr-arr[curr] >=0)
                if (visited.getOrDefault(curr-arr[curr], false) == false) {
                    visited.put(curr-arr[curr], true);
                    queue.offer(curr-arr[curr]);
                }

            if (curr+arr[curr] < size)
                if (visited.getOrDefault(curr+arr[curr], false) == false) {
                    visited.put(curr+arr[curr], true);
                    queue.offer(curr+arr[curr]);
                }
        }

        return false;
    }
}
