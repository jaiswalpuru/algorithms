class Solution {
public:
    int maxWidthRamp(vector<int>& nums) {
        int max_width = 0;
        stack<int> st;
        int ind = -1;

        for (int i = 0; i < nums.size(); i++) {
            if (st.empty() || nums[st.top()] > nums[i]) {
                st.push(i);
            }
        }
        
        for (int j = nums.size() - 1; j >= 0; j--) {
            while (!st.empty() && nums[st.top()] <= nums[j]) {
                max_width = max(max_width, j - st.top());
                st.pop();
            }
        }

        return max_width;
    }
};
