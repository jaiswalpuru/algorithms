class Solution {
    int res;
    int size;
    int diff;
    public int threeSumClosest(int[] nums, int target) {
        res = 0;
        size = nums.length;
        diff = Integer.MAX_VALUE; 
        Arrays.sort(nums);

        for (int i=0; i<size-2; i++) {
            if (i == 0 || nums[i] != nums[i-1]) {
                twoSumClosest(nums, target, i);
            }
        }

        return res;
    }

    private void twoSumClosest(int[] nums, int target, int ind) {
        int low = ind+1;
        int high = size-1;

        while(low < high) {
            int sum = nums[ind] + nums[low] + nums[high];
            int currDiff = Math.abs(target - sum);
            if (diff >= currDiff) {
                diff = currDiff;
                res = sum;      
            }
            if (sum == target) {
                low++;
                high--;
                while(low<high && nums[low] == nums[low-1]) low++;
            }else if (sum > target) {
                high--;
            }else {
                low++;
            }
        }
    }
}
