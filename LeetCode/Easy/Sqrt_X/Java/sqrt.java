class Solution {
    public int mySqrt(int x) {
        int l = 0;
        int h = x;
        long num;
        while (l <= h) {
            int mid = l + (h-l)/2;
            num = (long)mid*mid;
            if (num < x) {
                l = mid+1;
            }else if (num > x){
                h = mid-1;
            }else {
                return mid;
            }
        }
        return h;
    }
}

