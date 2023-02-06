class Solution {
    public int[] shuffle(int[] nums, int n) {
        int size = 2*n;
        int i=0, j=n;

        int[] res = new int[size];
        int k = 0;
        while(i < n && j < size) {
            res[k] = nums[i++];
            res[k+1] = nums[j++];
            k+=2;
        }
        return res;
    }
}
