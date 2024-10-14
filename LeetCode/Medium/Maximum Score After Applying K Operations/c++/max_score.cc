class Solution {
public:
    long long maxKelements(vector<int>& nums, int k) {
        int val = 0;
        long long sum = 0;
        priority_queue<int> pq;

        for (int num : nums) pq.push(num);

        while (!pq.empty() && k > 0) {
            val = pq.top();
            pq.pop();
        
            sum += val;
            pq.push(ceil(val / 3.0));
            k--;
        }

        return sum;
    }
};
