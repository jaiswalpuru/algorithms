class Solution {
    public int minimizeMax(int[] nums, int p) {
        Arrays.sort(nums);
        int n = nums.length;
        int l=0, h=nums[n-1]-nums[0];
        while(l < h) {
            int mid = l+(h-l)/2;
            if (countValidPairs(nums, mid) >= p) h = mid;
            else l = mid+1;
        }
        return l;
    }

    private int countValidPairs(int[] nums, int thres) {
        int ind=0, cnt=0;
        while(ind < nums.length-1) {
            if (nums[ind+1] - nums[ind] <= thres) {
                cnt++;
                ind++;
            }
            ind++;
        }
        return cnt;
    }
}
