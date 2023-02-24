class Solution {
    public int minimumDeviation(int[] nums) {
        int size = nums.length;
        int minEvenElement = Integer.MAX_VALUE;
        int minDeviation = Integer.MAX_VALUE;
        PriorityQueue<Integer> pq = new PriorityQueue<>(size, Collections.reverseOrder());

        // push all the elements in the priority queue
        // if odd multiply them with 2 and then push
        for (Integer val : nums) {
            if (val%2 == 0) {
                minEvenElement = Math.min(minEvenElement, val);
                pq.offer(val);
            }else {
                minEvenElement = Math.min(minEvenElement, val*2);
                pq.offer(val*2);
            }
        }

        // find minimum deviation
        while(pq.size() > 0) {
            int topElement = pq.poll();
            minDeviation = Math.min(minDeviation, topElement-minEvenElement);
            if (topElement % 2 == 0) {
                minEvenElement = Math.min(minEvenElement, topElement/2);
                pq.offer(topElement/2);
            }else {
                break;
            }
        }

        return minDeviation;
    }
}
