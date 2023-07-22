// method 1 
class Solution {
    public boolean isPowerOfTwo(int n) {
        if (n == 0) return false;
        while(n%2 == 0) n = n/2;
        return n == 1;
    }
}

// method 2 without loop or recursion
class Solution {
    public boolean isPowerOfTwo(int n) {
        if (n == 0) return false;
        return ((long)n&((long)n-1)) == 0;
    }
}
