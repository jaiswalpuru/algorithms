//method 1 using hashmap
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

//method 2 using bitmasking
class Solution {
    public int[] singleNumber(int[] nums) {
        int bitMask = 0;
        for (int num : nums) bitMask ^= num;
        int diff = bitMask&(-bitMask);
        int x = 0;
        for (int num : nums) if ((num&diff) != 0) x ^= num;
        return new int[]{x, bitMask^x};
    }
}
