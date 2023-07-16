/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* productExceptSelf(int* nums, int numsSize, int* returnSize){
    int* pre = (int*)malloc(numsSize*sizeof(int));
    int* suff = (int*)malloc(numsSize*sizeof(int));
    pre[0] = 1;
    suff[numsSize-1] = 1;
    for (int i=1; i<numsSize; i++) *(pre+i) = (*(pre+i-1)) * (*(nums+i-1));
    for (int i=numsSize-2; i>=0; i--) *(suff+i) = (*(suff+i+1)) * (*(nums+i+1));
    int* res = (int*) malloc(numsSize*sizeof(int));
    for (int i=0; i<numsSize; i++) *(res+i) = (*(pre+i)) * (*(suff+i));
    *returnSize = numsSize;
    return res;
}
