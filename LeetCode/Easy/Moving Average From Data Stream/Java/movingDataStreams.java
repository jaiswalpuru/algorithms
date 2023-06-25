class MovingAverage {
    ArrayDeque<Integer> dq;
    int size;
    int sum = 0;
    public MovingAverage(int size) {
        this.size = size;
        dq = new ArrayDeque<>();
    }
    
    public double next(int val) {
        if (dq.size() == size) sum -= dq.removeFirst();
        dq.add(val);
        sum += val;
        return (double)sum/dq.size();
    }
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * MovingAverage obj = new MovingAverage(size);
 * double param_1 = obj.next(val);
 */
