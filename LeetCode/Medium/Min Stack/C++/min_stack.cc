class MinStack {
public:
    stack<pair<int,int> > st_pair;
    MinStack() {}
    
    void push(int val) {
        int min_ele;
        if (st_pair.empty()) min_ele = val;
        else min_ele = min(val, st_pair.top().second);
        st_pair.push(make_pair(val, min_ele));
    }
    
    void pop() {
        st_pair.pop();
    }
    
    int top() {
        return st_pair.top().first;
    }
    
    int getMin() {
        return st_pair.top().second;
    }
};

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack* obj = new MinStack();
 * obj->push(val);
 * obj->pop();
 * int param_3 = obj->top();
 * int param_4 = obj->getMin();
 */
