class Solution {
public:
    int minSwaps(vector<int>& nums) {
        int grp_zeroes = min_swap_helper(nums, 0);
        int grp_ones = min_swap_helper(nums, 1);
        return min(grp_zeroes, grp_ones);
    }

    int min_swap_helper(vector<int> &nums, int val) {
        int len = nums.size();
        int val_cnt = 0;
        int st = 0, en = 0;
        int max_val_in_win = 0;
        int curr_val_in_win = 0;

        for (int i = 0; i < len; i++) if (val == nums[i]) val_cnt++;
        if (val_cnt == 0 || val_cnt == len) return 0;

        while (en < val_cnt) if (nums[en++] == val) curr_val_in_win++;
        max_val_in_win = curr_val_in_win;

        while (en < len) {
            if (nums[st++] == val) curr_val_in_win--;
            if (nums[en++] == val) curr_val_in_win++;
            max_val_in_win = max(max_val_in_win, curr_val_in_win);
        }
            
        return val_cnt - max_val_in_win;
    }

};
