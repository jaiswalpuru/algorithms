class Solution {
    public boolean checkIfPangram(String sentence) {
        int[] chars = new int[26];
        Arrays.fill(chars, -1);
        for (int i=0; i<sentence.length(); i++)chars[sentence.charAt(i)-'a'] = 1;
        for (int i=0; i<26; i++) {
            if (chars[i] == -1) return false;
        }
        return true;
    }
}
