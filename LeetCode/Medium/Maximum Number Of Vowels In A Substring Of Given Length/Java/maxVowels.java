class Solution {
    public int maxVowels(String s, int k) {
        int cnt = 0, max = 0;
        int i = 0, j = 0;
        while(j < s.length()) {
            char c = s.charAt(j);
            if (j-i == k) {
                max = Math.max(max, cnt);
                if (isVowel(s.charAt(i))) cnt--;
                i++;
            }
            if (isVowel(s.charAt(j))) {
                cnt++;
            }
            j++;
        }
        max = Math.max(max, cnt);
        return max;
    }

    private boolean isVowel(char c) {
        return c == 'a' ||
            c == 'e' ||
            c == 'i' ||
            c == 'o' ||
            c == 'u';
    }
}
