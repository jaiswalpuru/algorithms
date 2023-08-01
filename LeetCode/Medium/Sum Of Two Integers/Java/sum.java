class Solution {
    public int getSum(int a, int b) {
        int x=Math.abs(a), y=Math.abs(b);
        if (x < y) return getSum(b, a);
        int sign = a > 0 ? 1 : -1;

        if (a*b >= 0) {
            while(y != 0) {
                int ans = x^y;
                int carry = (x&y) << 1;
                x = ans;
                y = carry;
            }
        } else {
            while(y != 0) {
                int ans = x^y;
                int carry = ((~x)&y) << 1;
                x = ans;
                y = carry;
            }
        }
        return sign*x;
    }
}
