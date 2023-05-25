class Solution {
    public int reverse(int x) {
        boolean isNegative = x < 0;
        if (isNegative) x = -x;
        long reverseNumber = 0;
        while(x > 0) {
            reverseNumber = (reverseNumber*10) + (x%10);
            x = x/10;
        }
        if (isNegative) {
            reverseNumber = -reverseNumber;
            return reverseNumber < Integer.MIN_VALUE ? 0 : (int)reverseNumber;
        }
        return reverseNumber > Integer.MAX_VALUE ? 0 : (int)reverseNumber;
    }
}
