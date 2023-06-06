class Solution {
    public int[] singleNumber(int[] nums) {
        HashMap<Integer, Integer> cntMap = new HashMap<>();
        for (int val : nums) cntMap.put(val, cntMap.getOrDefault(val, 0) + 1);
        int[] output = new int[2];
        int i = 0;
        for (Integer k : cntMap.keySet()) {
            if (cntMap.get(k) == 1) output[i++] = k;
        }
        return output;
    }
}
