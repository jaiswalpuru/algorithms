// One brute force approach would be sort an array and then just
// iterate and find the dup

class Solution {
    public int findDuplicate(int[] nums) {
        int size = nums.length;
        int low = 0;
        int high = size;
        int dup = -1;

        while(low <= high) {
            int mid = low + (high-low)/2;

            int cnt = 0;
            for (int val : nums)
                if (val <= mid)
                    cnt++;

            if (cnt > mid) {
                dup = mid;
                high = mid-1;
            }else {
                low = mid+1;
            }
        }
        return dup;
    }
}
