/**
 * Note: The returned array must be malloced, assume caller calls free().
 */

int count_ones(int n) {
    int cnt = 0;
    while(n > 0) {
        if ((n&1) == 1) cnt++;
        n >>= 1;
    }
    return cnt;
}

int* countBits(int n, int* returnSize){
    int* res = (int*) malloc(sizeof(int)*(n+1));
    for (int i=0; i<=n; i++) *(res+i) = count_ones(i);
    *returnSize = n+1;
    return res;
}
