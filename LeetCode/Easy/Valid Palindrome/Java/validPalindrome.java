class Solution {
    public boolean isPalindrome(String s) {
        s = s.toLowerCase();
        StringBuilder sb = new StringBuilder();
        for (int i=0; i<s.length(); i++)
            if (Character.isAlphabetic(s.charAt(i)) || Character.isDigit(s.charAt(i))) 
                sb.append(s.charAt(i));
                
        return sb.toString().equals(sb.reverse().toString());
    }
}
