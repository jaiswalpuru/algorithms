class Solution {
public:
    int search(vector<int>& nums, int target) {
        return _search(nums, target, 0, nums.size()-1);
    }

    int _search(vector<int>& nums, int target, int l, int h) {
        if (l > h) return -1;
        int mid = l + (h-l)/2;
        if (nums[mid] == target) return mid;
        if (nums[mid] >= nums[l]) {
            if (target >= nums[l] && target <= nums[mid]) return _search(nums, target, l, mid-1);
            return _search(nums, target, mid+1, h);
        }
        if (target >= nums[mid] && target <= nums[h]) return _search(nums, target, mid+1, h);
        return _search(nums, target, l, mid-1);
    }
};
