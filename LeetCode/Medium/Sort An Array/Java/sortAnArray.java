class Solution {
    public int[] sortArray(int[] nums) {
        int[] temp = new int[nums.length];
        mergeSort(nums, 0, nums.length-1, temp);
        return nums;
    }

    public void mergeSort(int[] arr, int l, int r, int[] temp) {
        if (l >= r) return;
        int mid = (l+r)/2;
        mergeSort(arr, l, mid, temp);
        mergeSort(arr, mid+1, r, temp);
        merge(arr, l, mid, r, temp);
    }

    public void merge(int[] arr, int l, int mid, int r, int[] temp) {
        int s1 = l;
        int s2 = mid+1;
        int n = mid-l+1;
        int m = r-mid;

        for (int k=0; k<n; k++) temp[s1+k] = arr[s1+k];
        for (int k=0; k<m; k++) temp[s2+k] = arr[s2+k];
        
        int i=0, j=0, k=l;
        while (i<n && j<m) {
            if (temp[s1+i] < temp[s2+j]) arr[k++] = temp[s1+i++];
            else arr[k++] = temp[s2+j++];
        }

        while(i<n) arr[k++] = temp[s1+i++];

        while(j<m) arr[k++] = temp[s2+j++];
    }
}
