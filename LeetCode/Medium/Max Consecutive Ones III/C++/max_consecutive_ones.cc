class Solution {
public:
    int longestOnes(vector<int>& nums, int k) {
        int i=0;
        int j=0;
        int curr_one = 0;
        int z = 0;
        int max_one = 0;
        int t = k;
        while (j < nums.size()) {
            if (nums[j] == 1) curr_one++;
            else z++;
            if (z > k) {
                while(z > k) {
                    if (nums[i] == 0) z--;
                    i++;
                }
            } 
            max_one = max(max_one, j-i+1);
            j++;
        }
        return max_one;
    }
};
