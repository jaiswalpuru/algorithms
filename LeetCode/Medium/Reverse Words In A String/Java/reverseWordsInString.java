class Solution {
    public String reverseWords(String s) {
        String[] str = s.split(" ");
        List<String> strArr = new ArrayList<>();
        for (int i=0; i<str.length; i++) 
            if (str[i].length() > 0) strArr.add(str[i]);
            
        Collections.reverse(strArr);
        return String.join(" ", strArr);
    }
}
