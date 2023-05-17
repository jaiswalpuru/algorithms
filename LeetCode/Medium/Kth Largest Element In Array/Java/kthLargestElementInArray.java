class Solution {
    public int findKthLargest(int[] nums, int k) {
        PriorityQueue<Integer> pq = new PriorityQueue<>(Collections.reverseOrder());
        for (int val : nums) pq.add(val);
        while(k-- > 1) pq.poll();
        return pq.peek();
    }
}
