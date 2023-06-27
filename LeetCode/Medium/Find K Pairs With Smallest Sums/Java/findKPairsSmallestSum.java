class Solution {
    public List<List<Integer>> kSmallestPairs(int[] nums1, int[] nums2, int k) {
        List<List<Integer>> res = new ArrayList<>();
        PriorityQueue<T> pq = new PriorityQueue<>((a, b) -> a.sum-b.sum);
        for (int i=0; i<nums1.length; i++) {
            pq.offer(new T(i, 0, nums1[i]+nums2[0]));
        }

        while(k > 0 && !pq.isEmpty()) {
            T curr = pq.poll();
            res.add(Arrays.asList(nums1[curr.x], nums2[curr.y]));
            if (curr.y + 1 < nums2.length) {
                pq.offer(new T(curr.x, curr.y+1, nums1[curr.x] + nums2[curr.y+1]));
            }
            k--;
        }
        return res;
    }
}

class T {
    int x, y, sum;
    public T(int x, int y, int sum) {
        this.x = x;
        this.y = y;
        this.sum = sum;
    }
}
