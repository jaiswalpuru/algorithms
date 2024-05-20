class Solution {
public:
    int findMaxK(vector<int>& nums) {
        int max_val = -1;
        map<int, bool> visited;

        for (auto num : nums) num < 0 ? visited[num] = true : visited[num] = false ;
        for (auto num : nums) {
            if (num > 0) {
                if (visited.find(-num) != visited.end()) {
                    max_val = max(max_val, num);
                }
            }
        }
        return max_val;
    }
};
