class Solution {
public:
    int count_valid(vector<int> nums, int mid) {
        int ind = 0;
        int cnt = 0;
        while (ind < nums.size()-1) {
            if (nums[ind+1]-nums[ind] <= mid) {
                cnt++;
                ind++;
            }
            ind++;
        }
        return cnt;
    }

    int minimizeMax(vector<int>& nums, int p) {
        sort(nums.begin(), nums.end());
        int l=0, h=nums[nums.size()-1]-nums[0];
        while(l < h) {
            int mid=l+(h-l)/2;
            if (count_valid(nums, mid) >= p) h = mid;
            else l=mid+1;
        }
        return l;
    }
};
