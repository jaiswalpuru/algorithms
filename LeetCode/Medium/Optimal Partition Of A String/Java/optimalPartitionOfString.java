class Solution {
    public int partitionString(String s) {
        int partition = 0;
        HashMap<Character, Integer> cntMap = new HashMap<>();
        for (int i=0; i<s.length(); i++) {
            int cnt = cntMap.getOrDefault(s.charAt(i), 0);
            if (cnt >= 1) {
                partition++;
                cntMap = new HashMap<>();
            }
            cntMap.put(s.charAt(i), 1);
        }
        return partition+1;
    }
}
