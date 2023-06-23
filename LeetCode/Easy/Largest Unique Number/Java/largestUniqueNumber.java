class Solution {
    public int largestUniqueNumber(int[] nums) {
        HashMap<Integer, Integer> cntMap = new HashMap<>();
        for (int val : nums) cntMap.put(val, cntMap.getOrDefault(val, 0)+1);
        int res = -1;
        for (Integer key : cntMap.keySet()) {
            if (cntMap.get(key) == 1) res = Math.max(res, key);
        }
        return res;
    }
}
