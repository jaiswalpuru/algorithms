// method 1
class Solution {
    public void rotate(int[] nums, int k) {
        int n = nums.length-1;
        int[] res = new int[n+1];
        if (n+1 <= k) k = k%(n+1);
        int j=0;
        for (int i=n-k+1; i<=n; i++) res[j++] = nums[i];
        for (int i=0; i<=n-k; i++) res[j++] = nums[i];
        System.arraycopy(res, 0, nums, 0, nums.length);
    }
}

//method 2
// first reverse the whole array
// then reverse the first k elements
// finally reverse the whole array
class Solution {
    public void rotate(int[] nums, int k) {
        int n = nums.length;
        k = k%n;
        reverse(nums, 0, nums.length-1);
        reverse(nums, 0, k-1);
        reverse(nums, k, n-1);
    }

    public void reverse(int[] nums, int start, int end) {
        while (start < end) {
            int temp = nums[start];
            nums[start] = nums[end];
            nums[end] = temp;
            start++;
            end--;
        }
    }
}
