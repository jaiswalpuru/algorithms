class Solution {
    public List<List<Integer>> minimumAbsDifference(int[] arr) {
        Arrays.sort(arr);
        int size = arr.length;
        List<List<Integer>> res = new ArrayList<>();
        
        int minDiff = Integer.MAX_VALUE;
        for (int i=0; i<size-1; i++) minDiff = Math.min(minDiff, arr[i+1]-arr[i]);

        for (int i=1; i<size; i++)
            if (arr[i]-arr[i-1] == minDiff) 
                res.add(Arrays.asList(arr[i-1], arr[i]));
        return res;
    }
}
