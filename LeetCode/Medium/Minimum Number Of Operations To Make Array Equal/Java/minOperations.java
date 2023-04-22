class Solution {
    public int minOperations(int n) {
        int[] arr = new int[n];
        for (int i=0; i<n; i++) {
            arr[i] = (2*i)+1;
        }

        int totalMoves = 0;
        int target = 0;
        
        for (int val : arr) target += val;
        target /= n;

        for (int val : arr) totalMoves += Math.abs(val-target);

        return totalMoves/2;
    }
}
