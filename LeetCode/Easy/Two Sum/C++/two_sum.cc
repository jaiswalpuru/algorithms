class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        std::map<int, int> cache;
        std::vector<int> ans;

        for (int i = 0; i < nums.size(); i++) cache[nums[i]] = i;
        for (int i = 0; i < nums.size(); i++) {
            int val = target-nums[i];
            if (cache.find(val) != cache.end()) {
                if (cache[val] == i) continue;
                ans.push_back(i);
                ans.push_back(cache[val]);
                break;
            }
        }
        return ans;
    }
};
