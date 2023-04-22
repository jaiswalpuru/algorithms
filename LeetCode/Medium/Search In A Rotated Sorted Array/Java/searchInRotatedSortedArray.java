class Solution {
    public int search(int[] nums, int target) {
        return searchUtil(nums, 0, nums.length-1, target);
    }
    
    public int searchUtil(int[] nums, int low, int high, int target) {
        if (low > high) {
            return -1;
        }
        int mid = low + (high-low)/2;
        if (nums[mid] == target) {
            return mid;
        }
        if (nums[mid] >= nums[low]) {
            if (nums[mid] > target && nums[low] <= target) {
                return searchUtil(nums, low, mid-1, target);    
            }
            return searchUtil(nums, mid+1, high, target);
        }
        if (target > nums[mid] && target <= nums[high]) {
            return searchUtil(nums, mid+1, high, target);
        }else {
            return searchUtil(nums, low, mid-1, target);
        }
    }
}
