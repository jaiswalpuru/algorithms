class Solution {
    public int lastStoneWeight(int[] stones) {
        PriorityQueue<Integer> pq = new PriorityQueue<>(Collections.reverseOrder());
        for (int val : stones) pq.add(val);
        while (pq.size() >= 2) {
            int s1 = pq.poll();
            int s2 = pq.poll();
            if (s1 != s2) pq.offer(Math.abs(s1-s2));
        }
        
        return pq.size() > 0 ? pq.peek() : 0;
    }
}
