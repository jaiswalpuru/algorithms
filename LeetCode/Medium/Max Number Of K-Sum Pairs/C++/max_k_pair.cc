class Solution {
public:
    int maxOperations(vector<int>& nums, int k) {
        sort(nums.begin(), nums.end());
        int i=0, j=nums.size()-1;
        int op = 0;
        while (i < j) {
            int sum = nums[i] + nums[j];
            if (sum == k) {
                i++;
                j--;
                op++;
            } else if (sum > k) {
                j--;
            } else {
                i++;
            }
        }
        return op;
    }
};
