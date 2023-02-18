class Solution {
    public int numPairsDivisibleBy60(int[] time) {
        int[] remainders = new int[60];
        int pairs = 0;

        for (int i=0; i<time.length; i++) {
            if (time[i]%60 == 0) pairs += remainders[0];
            else pairs += remainders[60-time[i]%60];
            remainders[time[i]%60]++;
        }
        return pairs;
    }
}
