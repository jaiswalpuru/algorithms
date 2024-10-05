class Solution {
public:
    int subarraysDivByK(vector<int>& nums, int k) {
        int res = 0;
        int sum = 0;
        unordered_map<int, int> ind;
        ind[0] = 1;

        for (int i = 0; i < nums.size(); i++) {
            sum = ((sum + nums[i]) % k + k) % k;
            res += ind[sum];
            ind[sum]++;
        }

        return res;
    }
};
