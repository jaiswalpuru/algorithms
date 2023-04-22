class Solution {
    int size;
    public List<List<Integer>> threeSum(int[] nums) {
        size = nums.length;
        Arrays.sort(nums);
        
        List<List<Integer>> res = new ArrayList<>();

        for (int i=0; i<size; i++) {
            if (i == 0 || nums[i] != nums[i-1]) {
                List<List<Integer>> r = twoSum(nums, i, 0);
                if (r.size()>0) res.addAll(r);
            }
        }

        return res;
    }

    private List<List<Integer>> twoSum(int []nums, int ind, int target) {
        int low = ind+1;
        int high = size-1;
        List<List<Integer>> res = new ArrayList<>();
        List<Integer> r = new ArrayList<>();

        while(low < high) {
            int sum = nums[low] + nums[high] + nums[ind];
            if (sum == 0) {
                r = new ArrayList<>();
                r.add(nums[ind]);
                r.add(nums[low]);
                r.add(nums[high]);
                low++;
                high--;
                res.add(r);
                while(low < high && nums[low] == nums[low-1]) low++;
            }else if (sum > 0) {
                high--;
            }else {
                low++;
            }
        }
        
        return res;
    }
}
