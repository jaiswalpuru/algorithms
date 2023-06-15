class Solution {
    public int[] dailyTemperatures(int[] temperatures) {
        int[] answer = new int[temperatures.length];
        ArrayDeque<Integer> stack = new ArrayDeque<>();
        int i = 0;
        stack.push(i);
        for (int j=1; j<temperatures.length; j++) {
            while(!stack.isEmpty() && temperatures[j] > temperatures[stack.peek()]) {
                int index = stack.pop();
                answer[index] = j - index;
            }
            stack.push(j);
        }
        return answer;
    }
}
