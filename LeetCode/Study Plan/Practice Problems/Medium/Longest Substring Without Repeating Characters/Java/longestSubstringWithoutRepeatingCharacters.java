class Solution {
    public int lengthOfLongestSubstring(String s) {
        int res = 0;
        int j=0;
        HashMap<Character, Integer> cnt = new HashMap<>();
        for (int i=0; i<s.length(); i++) {
            cnt.put(s.charAt(i), cnt.getOrDefault(s.charAt(i), 0)+1);
            while(cnt.get(s.charAt(i)) > 1) {
                cnt.put(s.charAt(j), cnt.get(s.charAt(j))-1);
                if (cnt.get(s.charAt(j)) == 0) cnt.remove(s.charAt(j));
                j++;
            }
            res = Math.max(res, i-j+1);
        }
        return res;
    }
}
