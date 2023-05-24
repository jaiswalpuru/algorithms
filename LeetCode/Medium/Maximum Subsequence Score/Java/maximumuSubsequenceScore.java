class P {
    int a, b;
    public P (int a, int b) {
        this.a = a;
        this.b = b;
    }
    @Override
    public String toString() { return a + "->" + b; } 
}

class Solution {
    public long maxScore(int[] nums1, int[] nums2, int k) {
        int n = nums1.length;
        PriorityQueue<Integer> pq = new PriorityQueue<>((x, y) -> x-y);
        P[] pair = new P[n];
        for (int i=0; i<n; i++)
            pair[i] = new P(nums1[i], nums2[i]);
        
        Arrays.sort(pair, (x, y) -> y.b-x.b);
        long sumK = 0;
        for (int i=0; i<k; i++) {
            sumK += pair[i].a;
            pq.add(pair[i].a);
        }

        long answer = sumK * pair[k-1].b;
        for (int i=k; i<n; i++) {
            sumK += pair[i].a - pq.poll();
            pq.add(pair[i].a);
            answer = Math.max(answer, sumK * pair[i].b);
        }
        return answer;
    }
}
