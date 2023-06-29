/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* findDuplicates(int* nums, int numsSize, int* returnSize){
    int ind = 0;
    int *res = (int *) malloc(numsSize * sizeof(int));
    int dup = 0;
    for (int i=0; i<numsSize; i++) {
        ind = abs(nums[i])-1;
        if (nums[ind] < 0)
            res[dup++] = abs(nums[i]);
        nums[ind] = -nums[ind];
    }
    *returnSize = dup;
    return res;
}
