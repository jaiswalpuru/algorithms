#define MAX(a,b) (((a) > (b)) ? (a) : (b))

int jump(int* nums, int numsSize){
    int curr_far = 0, curr_end = 0;
    int res = 0;
    for (int i=0; i<numsSize-1; i++) {
        curr_far = MAX(curr_far, i+*(nums+i));
        if (i == curr_end) {
            res++;
            curr_end = curr_far;
        }
    }
    return res;
}
