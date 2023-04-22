class Solution {
    public boolean checkInclusion(String s1, String s2) {
        int s1Size = s1.length();
        int s2Size = s2.length();

        if (s2Size < s1Size) return false;
        
        int[] s1CharCnt = new int[26];
        for (char c : s1.toCharArray()) {
            s1CharCnt[c-'a']++;
        }

        for (int i=0; i<=s2Size-s1Size; i++) {
            int []s2CharCnt = new int[26];
            for (int j = i; j<i+s1Size; j++) {
                s2CharCnt[s2.charAt(j)-'a']++;
            }
            if (match(s1CharCnt, s2CharCnt)) return true;
        }

        return false;
    }

    public boolean match(int[] s1, int[] s2) {
        for (int i=0; i<26; i++)
            if (s1[i] != s2[i]) return false;
        
        return true;
    }
}
