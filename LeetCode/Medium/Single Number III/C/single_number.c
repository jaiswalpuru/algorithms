/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* singleNumber(int* nums, int numsSize, int* returnSize){
    *returnSize = 2;
    int* res = (int *) malloc(sizeof(int)*2);
    int bit_mask = 0;
    for (int i=0; i<numsSize; i++) bit_mask ^= nums[i];
    long diff = (long)bit_mask&(-(long)bit_mask);
    int x = 0;
    for (int i=0; i<numsSize; i++) if (((long)nums[i]&diff) != 0) x ^= nums[i];
    res[0] = x;
    res[1] = bit_mask^x;
    return res;
}
