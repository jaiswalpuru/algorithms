class T {
    int k, v;
    public T(int k, int v) {
        this.k = k;
        this.v = v;
    } 
}

class Solution {
    public long totalCost(int[] costs, int k, int candidates) {
        PriorityQueue<T> pq = new PriorityQueue<>((a, b) -> {
            return a.k == b.k ? a.v-b.v : a.k-b.k;
        });
        int n = costs.length;
        for (int i=0; i<candidates; i++) pq.offer(new T(costs[i], 0));
        for (int i=Math.max(n-candidates, candidates); i<n; i++) pq.offer(new T(costs[i], 1));

        long totalCost = 0;
        int head = candidates;
        int tail = n-1-candidates;
        for (int i=0; i<k; i++) {
            T curr = pq.poll();
            totalCost += curr.k;
            if (head <= tail) {
                if (curr.v == 0) {
                    pq.offer(new T(costs[head], 0));
                    head++;
                } else {
                    pq.offer(new T(costs[tail], 1));
                    tail--;
                }
            }
        }

        return totalCost;
    }
}
