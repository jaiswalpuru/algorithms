class Solution {
public:
    vector<int> nextGreaterElement(vector<int>& nums1, vector<int>& nums2) {
        unordered_map<int, int> u_map;
        vector<int> res(nums1.size());

        for (int i = 0; i < nums2.size(); i++) {
            u_map[nums2[i]] = i;
        }

        for (int i = 0; i < nums1.size(); i++) {
            int ind = u_map[nums1[i]];
            int num = nums1[i];
            for (int j = ind + 1; j < nums2.size(); j++) {
                if (num < nums2[j]) {
                    num = nums2[j];
                    break;
                }
            }

            res[i] = (num == nums1[i] ? -1 : num);
        }
        
        return res;
    }
};
