class Solution {
    public boolean checkAlmostEquivalent(String word1, String word2) {
        int[] word1Char = new int[26];
        int[] word2Char = new int[26];
        int n = word1.length();
        for (int i=0; i<n; i++) {
            word1Char[word1.charAt(i)-'a']++;
            word2Char[word2.charAt(i)-'a']++;
        }
        for (int i=0; i<26; i++) 
            if (Math.abs(word1Char[i]-word2Char[i]) > 3) return false;
        return true;
    }
}
