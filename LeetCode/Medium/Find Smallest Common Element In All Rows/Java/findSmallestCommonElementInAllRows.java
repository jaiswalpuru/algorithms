class Solution {
    public int smallestCommonElement(int[][] mat) {
        int n = mat.length;
        int m = mat[0].length;
        Map<Integer, Integer> cntMap = new HashMap<>();
        for (int i=0; i<n; i++) {
            for (int j=0; j<m; j++) {
                cntMap.put(mat[i][j], cntMap.getOrDefault(mat[i][j], 0) + 1);
            }
        }

        int minVal = Integer.MAX_VALUE;
        for (Integer k : cntMap.keySet()) {
            if (cntMap.get(k) == n) minVal = Math.min(minVal, k);
        }

        return minVal == Integer.MAX_VALUE ? -1 : minVal;
    }
}
