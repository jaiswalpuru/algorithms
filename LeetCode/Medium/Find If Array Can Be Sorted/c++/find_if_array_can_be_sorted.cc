class Solution {
public:
    bool canSortArray(vector<int>& nums) {
        vector<int> vals = nums;

        for (int i = 0; i < vals.size(); i++) {
            for (int j = 0; j < vals.size() - i - 1; j++) {
                if (vals[j] > vals[j + 1]) {
                    if (__builtin_popcount(vals[j]) == __builtin_popcount(vals[j + 1])) swap(vals[j], vals[j + 1]);
                    else return false;
                }
            }
        }

        return true;
    }
};
