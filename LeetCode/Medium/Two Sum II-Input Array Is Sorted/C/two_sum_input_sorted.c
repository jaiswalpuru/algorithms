/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* twoSum(int* numbers, int numbersSize, int target, int* returnSize){
    *returnSize = 2;
    int* res = (int*) malloc((*returnSize)*sizeof(int));
    int i=0, j=numbersSize-1;
    while(i < j) {
        int sum = numbers[i] + numbers[j];
        if (sum == target) {
            res[0] = i+1;
            res[1] = j+1;
            break;
        } else if (sum > target) {
            j--;
        } else {
            i++;
        }
    }
    return res;
}
