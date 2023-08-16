class Solution {
public:
    vector<vector<int>> findDifference(vector<int>& nums1, vector<int>& nums2) {
        map<int, bool> num1;
        map<int, bool> num2;
        for (int i=0; i<nums1.size(); i++) num1[nums1[i]] = true;
        for (int i=0; i<nums2.size(); i++) num2[nums2[i]] = true;
        vector<vector<int>> res;
        set<int> v1;
        set<int> v2;
        for (int i=0; i<nums1.size(); i++) {
            auto it = num2.find(nums1[i]);
            if (it == num2.end()) {
                v1.insert(nums1[i]);
            }
        }
        for (int i=0; i<nums2.size(); i++) {
            auto it = num1.find(nums2[i]);
            if (it == num1.end()) {
                v2.insert(nums2[i]);
            }
        }
        res.push_back(vector<int>(v1.begin(), v1.end()));
        res.push_back(vector<int>(v2.begin(), v2.end()));
        return res;
    }
};
