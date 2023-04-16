class Solution {
    public int minSteps(String s, String t) {
        HashMap<Character, Integer> sCnt = new HashMap<>();
        HashMap<Character, Integer> tCnt = new HashMap<>();
        for (char c : s.toCharArray()) sCnt.put(c, sCnt.getOrDefault(c, 0)+1);
        for (char c : t.toCharArray()) tCnt.put(c, tCnt.getOrDefault(c, 0)+1);

        int res = 0;
        for (Map.Entry<Character, Integer> m : sCnt.entrySet()) {
            int val = tCnt.getOrDefault(m.getKey(), 0);
            if (val < m.getValue()) {
                res += (m.getValue()-val);
            }
        }

        return res;
    }
}
