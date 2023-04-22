class Solution {
    public boolean validateStackSequences(int[] pushed, int[] popped) {
        Stack<Integer> stack = new Stack<>();
        int size = pushed.length;
        int k=0;

        for (int i=0; i<size; i++) {
            stack.push(pushed[i]);
            while (!stack.isEmpty() && stack.peek() == popped[k]) {
                stack.pop();
                k++;
            } 
        }

        return stack.isEmpty();
    }
}
