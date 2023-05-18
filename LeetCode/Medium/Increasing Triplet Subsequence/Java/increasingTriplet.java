class Solution {
    public boolean increasingTriplet(int[] nums) {
        int minOne = Integer.MAX_VALUE;
        int minTwo = Integer.MAX_VALUE;
        for (int val : nums) {
            if (minOne >= val) {
                minOne = val;
            } else if (minTwo >= val) {
                minTwo = val;
            } else {
                return true;
            }
        }
        return false;
    }
}
