int findDuplicate(int* nums, int numsSize){
    int low=1, high=numsSize-1;
    int dup = 0;
    while(low <= high) {
        int mid = low+(high-low)/2;
        int cnt = 0;
        for (int i=0; i<numsSize; i++)
            if(nums[i] <= mid) cnt++;
        
        if (cnt > mid) {
            dup = mid;
            high = mid-1;
        } else {
            low = mid+1;
        }
    }
    return dup;
}
