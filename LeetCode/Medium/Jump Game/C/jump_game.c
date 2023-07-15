#define MAX(x, y) (((x) > (y)) ? (x) : (y))

bool canJump(int* nums, int numsSize){
    int max_jump = nums[0];
    for (int i=1; i<numsSize; i++) {
        if (max_jump-1 < 0) return false;
        max_jump = MAX(max_jump-1, nums[i]);
    }
    return true;
}
