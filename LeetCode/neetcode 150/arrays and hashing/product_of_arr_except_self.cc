class Solution {
public:
    vector<int> productExceptSelf(vector<int>& nums) {
        int first_ele = nums[0];
        nums[0] = 1;
        int n = nums.size();
        vector<int> res(n, 0);
        res[n - 1] = 1;
 
        for (int i = n - 2; i >= 0; i--) res[i] = res[i + 1] * nums[i + 1];
        int curr_ele = first_ele;
        for (int i = 1; i < n; i++) {
            int t = nums[i];
            nums[i] = curr_ele * nums[i - 1];
            curr_ele = t;
        }
        for (int i = 0; i < n; i++) res[i] = res[i] * nums[i];
        return res;
    }
};
