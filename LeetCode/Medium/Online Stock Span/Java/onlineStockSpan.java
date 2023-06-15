class StockSpanner {
    HashMap<Integer, Integer> spanMap;
    ArrayDeque<Integer> stack;
    public StockSpanner() {
        spanMap = new HashMap<>();
        stack = new ArrayDeque<>();
    }
    
    public int next(int price) {
        if (stack.size() == 0) {
            spanMap.put(price, 1);
            stack.push(price);
        } else {
            int days = 1;
            while(!stack.isEmpty() && stack.peek() <= price) {
                days += spanMap.get(stack.pop());
            }
            spanMap.put(price, days);
            stack.push(price);
        }
        return spanMap.get(price);
    }
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * StockSpanner obj = new StockSpanner();
 * int param_1 = obj.next(price);
 */
