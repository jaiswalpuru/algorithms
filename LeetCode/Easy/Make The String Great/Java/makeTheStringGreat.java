class Solution {
    public String makeGood(String s) {
        StringBuilder sb = new StringBuilder(s);
        StringBuilder res = new StringBuilder();
        ArrayDeque<Character> st = new ArrayDeque<>();
        int i = 0;
        while(i < sb.length()) {
            Character c = sb.charAt(i);
            while (!st.isEmpty()) {
                Character t = st.peek();
                if (Character.toLowerCase(t) == Character.toLowerCase(c)) {
                    if ((Character.isLowerCase(t) && Character.isUpperCase(c))||
                    Character.isUpperCase(t) && Character.isLowerCase(c)) {
                        st.pop();
                        c = ' ';
                    }
                    break;
                } else break;
            }
            if (c != ' ')
                st.push(c);
            i++;
        }
        while(!st.isEmpty()) res.append(st.pop());
        return res.reverse().toString();
    }
}
