class MyStack {
    private Queue<Integer> q1;
    private Queue<Integer> q2;
    int top;
    public MyStack() {
        q1 = new LinkedList<>();
        q2 = new LinkedList<>();
    }
    
    public void push(int x) {
        q1.add(x);
        top = x;
    }
    
    public int pop() {
        while (q1.size() > 1) {
            top = q1.remove();
            q2.add(top);
        }
        int val = q1.remove();
        Queue<Integer> t = q1;
        q1 = q2;
        q2 = t;
        return val;
    }
    
    public int top() {
        return top;
    }
    
    public boolean empty() {
        return q1.size() == 0; 
    }
}

/**
 * Your MyStack object will be instantiated and called as such:
 * MyStack obj = new MyStack();
 * obj.push(x);
 * int param_2 = obj.pop();
 * int param_3 = obj.top();
 * boolean param_4 = obj.empty();
 */
