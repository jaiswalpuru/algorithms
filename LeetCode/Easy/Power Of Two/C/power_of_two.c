bool isPowerOfTwo(int n){
    if (n == 0) return false;
    return ((long)n&((long)n-1)) == 0;
}
