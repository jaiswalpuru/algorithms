class Solution {
public:
    vector<int> intersect(vector<int>& nums1, vector<int>& nums2) {
        map<int, int> mp1, mp2;
        vector<int> res;

        for (auto num : nums1) mp1[num]++;
        for (auto num : nums2) mp2[num]++;

        for (auto const &x : mp1) {
            int times = min(x.second, mp2[x.first]);
            for (int i = 0; i < times; i++) {
                res.push_back(x.first);
            }
        }
        
        return res;
    }
};
