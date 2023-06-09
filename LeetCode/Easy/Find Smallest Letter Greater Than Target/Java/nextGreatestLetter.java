class Solution {
    public char nextGreatestLetter(char[] letters, char target) {
        int[] chars = new int[26];
        Arrays.fill(chars, -1);
        for (int i=0; i<letters.length; i++)
            chars[letters[i]-'a'] = letters[i]-'a';
        int ind = target-'a'+1;
        for (int i=ind; i<26; i++) {
            if (chars[i] != -1) {
                return (char)(chars[i]+'a');
            }
        }
        return letters[0];
    }
}
