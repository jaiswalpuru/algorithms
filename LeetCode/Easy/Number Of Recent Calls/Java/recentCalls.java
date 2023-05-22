class RecentCounter {
    Queue<Integer> q;
    public RecentCounter() {
        q = new LinkedList<>();
    }
    
    public int ping(int t) {
        if (q.isEmpty()) {
            q.offer(t);
            return 1;
        }
        int low = t-3000;
        while(!q.isEmpty() && q.peek() < low) q.poll();
        q.offer(t);
        return q.size();
    }
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * RecentCounter obj = new RecentCounter();
 * int param_1 = obj.ping(t);
 */
