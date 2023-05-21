class Solution {
    public boolean uniqueOccurrences(int[] arr) {
        Map<Integer, Integer> valCnt = new HashMap<>();
        Map<Integer, Integer> occurCnt = new HashMap<>();
        for (int val : arr) valCnt.put(val, valCnt.getOrDefault(val, 0)+1);
        for (int k : valCnt.keySet()) {
            occurCnt.put(valCnt.get(k), occurCnt.getOrDefault(valCnt.get(k), 0)+1);
            if (occurCnt.get(valCnt.get(k)) > 1) return false;
        }
        return true;
    }
}
