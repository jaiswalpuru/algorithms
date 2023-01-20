class Solution {
    public int findMin(int[] nums) {
        int low = 0, high = nums.length-1;
        while(low < high) {
            int mid = low + (high-low)/2;
            if (nums[low] <= nums[mid]) {
                if (nums[low] > nums[high]) {
                    low = mid+1;
                }else {
                    high = mid;
                }
            }else {
                if (nums[low] > nums[high]) {
                    high = mid;
                }else {
                    low = mid;
                }
            }
        }
        return nums[low];
    }

