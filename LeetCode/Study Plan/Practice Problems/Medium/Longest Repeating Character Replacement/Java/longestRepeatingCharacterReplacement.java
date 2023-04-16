class Solution {
    public int characterReplacement(String s, int k) {
        HashMap<Character, Integer> freq = new HashMap<>();
        int maxFreq = 0;
        int res = 0;
        int j = 0;
        for (int i=0; i<s.length(); i++) {
            freq.put(s.charAt(i), freq.getOrDefault(s.charAt(i), 0)+1);
            maxFreq = Math.max(maxFreq, freq.get(s.charAt(i)));
            while(i-j-maxFreq+1 > k) {
                freq.put(s.charAt(j), freq.get(s.charAt(j))-1);
                j++;
            }
            res = Math.max(res, i-j+1);
        }
        return res;
    }
}
