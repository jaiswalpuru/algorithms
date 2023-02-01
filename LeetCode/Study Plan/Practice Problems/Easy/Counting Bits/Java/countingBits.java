class Solution {
    public int[] countBits(int n) {
        int[] ans = new int[n+1];
        ans[0] = 0;
        for (int i=1; i<=n; i++) {
            ans[i] = numberOfOneBits(i);
        }
        return ans;
    }

    public int numberOfOneBits(int n) {
        int oneBit = 0;
        int mask = 1;
        for (int i=0; i<32; i++) {
            if ((n&mask) != 0) {
                oneBit++;
            }
            mask <<= 1;
        }
        return oneBit;
    }
}
