class Solution {
public:
    vector<string> summaryRanges(vector<int>& nums) {
        string str = "";
        vector<string> res;
        if (nums.size() == 0) return res;
        int i=0, j=0;
        while(j < nums.size()-1) {
            if (((long)nums[j+1]-(long)nums[j]) == 1) {
                j++;
            } else {
                if (i != j) res.push_back(to_string(nums[i])+ "->" + to_string(nums[j]));
                else res.push_back(to_string(nums[j]));
                i=j+1;
                j++;
            }
        }
        if (i< j) res.push_back(to_string(nums[i]) + "->" + to_string(nums[j]));
        else if (j == nums.size()-1) res.push_back(to_string(nums[j]));
        return res;
    }
};
