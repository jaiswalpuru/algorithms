class Solution {
    public int maxDistance(List<List<Integer>> arrays) {
        int res = 0;
        int minVal = arrays.get(0).get(0);
        int maxVal = arrays.get(0).get(arrays.get(0).size()-1);
        for (int i=1; i<arrays.size(); i++) {
            int n = arrays.get(i).size();
            res = Math.max(res, 
                Math.max(
                    Math.abs(arrays.get(i).get(n-1)-minVal),
                    Math.abs(arrays.get(i).get(0)-maxVal)
                )
            );
            minVal = Math.min(minVal, arrays.get(i).get(0));
            maxVal = Math.max(maxVal, arrays.get(i).get(n-1));
        }
        return res;
    }
}
