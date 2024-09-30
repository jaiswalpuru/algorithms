class CustomStack {
    int max_size;
    vector<int> st;
    int i;
public:
    CustomStack(int maxSize) {
        i = -1;
        max_size = maxSize;
        st = vector<int>(max_size);
    }
    
    void push(int x) {
        if (i < max_size-1) st[++i] = x;
    }
    
    int pop() {
        if (i == -1) return -1;
        int res = st[i--];
        return res;
    }
    
    void increment(int k, int val) {
        for (int j = 0; j <= min(k-1, i); j++) st[j] = st[j] + val;
    }
};

/**
 * Your CustomStack object will be instantiated and called as such:
 * CustomStack* obj = new CustomStack(maxSize);
 * obj->push(x);
 * int param_2 = obj->pop();
 * obj->increment(k,val);
 */
