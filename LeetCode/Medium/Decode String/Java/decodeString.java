class Solution {
    public String decodeString(String s) {
        Stack<Character> stack = new Stack<>();
        StringBuilder sb = new StringBuilder();
        for (int i=0 ; i<s.length(); i++) {
            if (s.charAt(i) == ']') {
                StringBuilder temp = new StringBuilder();
                while (!stack.isEmpty() && stack.peek() != '[') temp.append(stack.pop());
                temp.reverse();
                stack.pop(); // to pop [
                String val = "";
                while (!stack.isEmpty() && Character.isDigit(stack.peek())) val = stack.pop() + val;
                
                if (val != "") {
                    Integer num = Integer.parseInt(val);
                    String t = temp.toString();
                    for (int k=0; k<num-1; k++) temp.append(t);
                }
                for (int k=0; k<temp.length(); k++) stack.push(temp.charAt(k));
            } else {
                stack.push(s.charAt(i));
            }
        }
        while(!stack.isEmpty()) sb.append(stack.pop());
        sb.reverse();
        return sb.toString();
    }
}
