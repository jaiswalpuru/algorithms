class Solution {
    public List<Integer> findDuplicates(int[] nums) {
        int n = nums.length;
        int i=0;

        while(i < n) {
            int ind = nums[i]-1;
            if (nums[i] != nums[ind]) {
                int temp = nums[i];
                nums[i] = nums[ind];
                nums[ind] = temp;
            }else {
                i++;
            }
        }

        List<Integer> res = new ArrayList<>();
        for (i=0; i<n; i++) {
            if (nums[i] != i+1) res.add(nums[i]);
        }
        return res;
    }
}
