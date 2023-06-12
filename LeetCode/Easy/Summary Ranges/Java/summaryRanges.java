class Solution {
    public List<String> summaryRanges(int[] nums) {
        List<String> res = new ArrayList<>();
        if (nums.length == 0) return res;
        int i = 0;
        int num = 0;
        for (i=1; i<nums.length; i++) {
            if (nums[i]-nums[i-1] != 1) {
                if (i-1 == num) res.add(Integer.toString(nums[num]));
                else res.add(nums[num]+"->"+nums[i-1]);
                num = i;
            }
        }
        if (i-num == 1) res.add(Integer.toString(nums[num]));
        else res.add(nums[num] + "->" + nums[i-1]);
        return res;
    }
}
