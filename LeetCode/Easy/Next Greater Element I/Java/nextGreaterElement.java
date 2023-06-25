class Solution {
    public int[] nextGreaterElement(int[] nums1, int[] nums2) {
        HashMap<Integer, Integer> index = new HashMap<>();
        for (int i=0; i<nums2.length; i++) index.put(nums2[i], i);
        int[] res = new int[nums1.length];
        for (int i=0; i<nums1.length; i++) {
            int ind = index.get(nums1[i]);
            int val = -1;
            for (int j=ind; j<nums2.length; j++) {
                if (nums1[i] < nums2[j]) {
                    val = nums2[j];
                    break;
                }
            }
            res[i] = val;
        }
        return res;
    }
}
