typedef struct Pair {
    int val;
    int ind;
} Pair;

int cmp(const void* a, const void* b) {
    const Pair* obj1 = a;
    const Pair* obj2 = b;
    return obj1->val-obj2->val;
}

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* twoSum(int* nums, int numsSize, int target, int* returnSize){
    int* res = (int*)malloc(sizeof(int)*2);
    Pair* p = (Pair*) malloc(sizeof(Pair)*numsSize);
    for (int i=0; i<numsSize; i++) {
        p[i].val = nums[i];
        p[i].ind = i;
    }
    qsort(p, numsSize, sizeof(Pair), cmp);
    int i=0, j=numsSize-1;
    while(i < j) {
        int sum = p[i].val + p[j].val;
        if (sum == target) {
            res[0] = p[i].ind;
            res[1] = p[j].ind;
            break;
        } else if (sum > target) {
            j--;
        } else {
            i++;
        }
    }
    *returnSize = 2;
    return res;
}
