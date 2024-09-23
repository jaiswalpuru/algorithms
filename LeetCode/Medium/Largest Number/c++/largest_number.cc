class Solution {
public:
    string largestNumber(vector<int>& nums) {
        string res = "";
        vector<string> nums_str;
        
        for (int i = 0; i < nums.size(); i++) nums_str.push_back(to_string(nums[i]));
        sort(nums_str.begin(), nums_str.end(), cmp);
        for (int i = 0; i < nums_str.size(); i++) res += nums_str[i];
        return res[0] == '0' ? "0" : res;
    }

    static bool cmp(const string &s1, const string &s2) {
        return (s1 + s2) > (s2 + s1) ? 1 : 0;
    }
};
