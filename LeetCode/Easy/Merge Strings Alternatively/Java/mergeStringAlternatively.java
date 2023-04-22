class Solution {
    public String mergeAlternately(String word1, String word2) {
        int i=0, j=0;
        boolean odd = false;
        StringBuilder sb = new StringBuilder();

        while (i<word1.length() && j<word2.length()) {
            if (!odd) {
                sb.append(word1.charAt(i));
                i++;
            } else {
                sb.append(word2.charAt(j));
                j++;
            }
            odd = !odd;
        }
        
        while(i < word1.length())
            sb.append(word1.charAt(i++));
        
        while(j < word2.length())
            sb.append(word2.charAt(j++));

        return sb.toString();
    }
}
