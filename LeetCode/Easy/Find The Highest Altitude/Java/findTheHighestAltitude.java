class Solution {
    public int largestAltitude(int[] gain) {
        int n = gain.length;
        int curr = Math.max(gain[0], 0);
        for (int i=1; i < n; i++) {
            gain[i] += gain[i-1];
            curr = Math.max(curr, gain[i]);
        }
        return curr;
    }
}
