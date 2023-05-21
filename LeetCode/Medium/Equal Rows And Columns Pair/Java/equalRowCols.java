class Solution {
    public int equalPairs(int[][] grid) {
        HashMap<String, Integer> row = new HashMap<>();
        int n = grid.length;

        int c = 0;
        for (int i=0; i<n; i++) {
            String s = Arrays.toString(grid[i]);
            row.put(s, row.getOrDefault(s, 0)+1);    
        }

        for (int j=0; j<n; j++) {
            int[] colArr = new int[n];
            for (int i=0; i<n; i++)
                colArr[i] = grid[i][j];

            c += row.getOrDefault(Arrays.toString(colArr), 0);
        }

        return c;
    }
}
