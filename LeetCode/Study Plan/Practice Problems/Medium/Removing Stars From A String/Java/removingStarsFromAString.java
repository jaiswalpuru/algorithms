class Solution {
    public String removeStars(String s) {
        Stack<Character> st = new Stack<>();
        StringBuilder sb = new StringBuilder();
        int size = s.length();
        int i=0;

        while(i < size) {
            if (s.charAt(i) == '*') {
                int cnt = 0;
                while(i < size && s.charAt(i) == '*') {
                    cnt++;
                    i++;
                }
                while(cnt-- > 0) st.pop();
            } else {
                st.push(s.charAt(i));
                i++;
            }
        }

        while(!st.isEmpty()) sb.append(st.pop());
        
        return sb.reverse().toString(); 
    }
}
