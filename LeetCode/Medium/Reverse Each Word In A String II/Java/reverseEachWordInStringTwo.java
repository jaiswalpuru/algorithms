class Solution {
    public void reverseWords(char[] s) {
        // reverse the whole string and then reverse the word.
        int n = s.length;
        reverse(s, 0, n-1);
        int i = 0, j = 0;
        //reverse each word
        while(i < n) {
            while(j < n && s[j] != ' ') j++;
            reverse(s, i, j-1);
            i = j+1;
            j++;
        }
    }

    private void reverse(char[] s, int i, int j) {
        while (i<j) {
            char t = s[i];
            s[i] = s[j];
            s[j] = t;
            i++;
            j--;
        }
    }
}
