class Solution {
public:
    vector<bool> kidsWithCandies(vector<int>& candies, int extraCandies) {
        int max_val = *max_element(candies.begin(), candies.end());
        vector<bool> res;
        for (int i=0; i<candies.size(); i++) {
            if (extraCandies + candies[i] >= max_val) res.push_back(true);
            else res.push_back(false);
        }
        return res;
    }
};
