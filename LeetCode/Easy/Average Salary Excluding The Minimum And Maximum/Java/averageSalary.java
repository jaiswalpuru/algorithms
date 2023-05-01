class Solution {
    public double average(int[] salary) {
        int max = -1;
        int min = (int)1e9;
        for (int val : salary) {
            max = Math.max(max, val);
            min = Math.min(min, val);
        }
        Integer total = Arrays.stream(salary).sum();
        total -= (max+min);
        return (double) total/(salary.length-2);
    }
}
