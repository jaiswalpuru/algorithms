class Solution {
    public int[] anagramMappings(int[] nums1, int[] nums2) {
        int size = nums1.length;
        Map<Integer, Integer> m2 = new HashMap<>();
        for (int i=0; i<size; i++) {
            m2.put(nums2[i], i);
        }

        int[] res = new int[nums1.length];
        for (int i=0; i<size; i++) {
            res[i] = m2.get(nums1[i]);
        }

        return res;
    }
}
