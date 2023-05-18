class Solution {
    public String reverseVowels(String s) {
        int i=0, j=s.length()-1;
        char[] sc = s.toCharArray();
        while(i < j) {
            while (i < j && !isVowel(sc[i])) i++;
            while (j > i && !isVowel(sc[j])) j--;
            if (i < j) {
                char c = sc[i];
                sc[i] = sc[j];
                sc[j] = c;
                i++;
                j--;
            }
        }
        return new String(sc);
    }

    private static boolean isVowel(char c) {
        return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || 
                c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U';
    }
}
