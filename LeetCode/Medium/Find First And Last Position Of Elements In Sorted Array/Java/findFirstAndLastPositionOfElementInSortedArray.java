class Solution {
    public int[] searchRange(int[] nums, int target) {
        int[] range = new int[2];
        range[0] = getLowerRange(nums, target);
        range[1] = getUpperRange(nums, target);
        return range;
    }
    
    public int getUpperRange(int[] nums, int target) {
        int pos = -1;
        if (nums == null || nums.length == 0) return pos;
        int low = 0;
        int high = nums.length-1;
        while(low <= high) {
            int mid = low + (high-low)/2;
            if (nums[mid] == target ) {
                pos = mid;
                low = mid+1;
            }else if (nums[mid] > target) {
                high = mid-1;
            }else {
                low = mid+1;
            }
        }
        return pos;
    }
    
    public int getLowerRange(int[] nums, int target) {
        int pos = -1;
        if (nums == null || nums.length == 0) return pos;
        int low = 0;
        int high = nums.length-1;
        while(low <= high) {
            int mid = low + (high-low)/2;
            if (nums[mid] == target ) {
                pos = mid;
                high = mid-1;
            }else if (nums[mid] > target) {
                high = mid-1; 
            }else { 
                low = mid+1;
            }
        }
        return pos;
    }
}
