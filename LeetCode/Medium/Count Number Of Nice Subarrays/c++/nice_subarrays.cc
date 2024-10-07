class Solution {
public:
    int numberOfSubarrays(vector<int>& nums, int k) {
        unordered_map<int, int> cnt_map;
        int res = 0;
        int sum = 0;
        cnt_map[sum] = 1;

        for (int i = 0; i < nums.size(); i++) {
            sum += nums[i] % 2;
            cout<<sum << " ";
            res += cnt_map[sum - k];
            cnt_map[sum]++;
        }

        return res;
    }
};
