class Solution {
    public int countElements(int[] arr) {
        HashMap<Integer, Integer> cntMap = new HashMap<>();
        
        for (int i=0; i<arr.length; i++)
            cntMap.put(arr[i], cntMap.getOrDefault(arr[i], 0)+1);
        
        int cnt = 0;
        for(int i=0; i<arr.length; i++)
            if (cntMap.containsKey(arr[i]+1)) cnt++;
        return cnt;
    }
}
