class Solution {
    public List<List<Integer>> findDifference(int[] nums1, int[] nums2) {
        Map<Integer, Boolean> m1 = new HashMap<>();
        Map<Integer, Boolean> m2 = new HashMap<>();
        List<List<Integer>> res = new ArrayList<>();

        for (int val : nums1) m1.put(val, true);
        for (int val : nums2) m2.put(val, true); 
        
        List<Integer> diff = new ArrayList<>();
        for (int val : m1.keySet()) if (!m2.containsKey(val)) diff.add(val);
        res.add(diff);
        
        diff = new ArrayList<>();
        for (int val : m2.keySet()) if (!m1.containsKey(val)) diff.add(val);

        res.add(diff);
        return res;
    }
}
