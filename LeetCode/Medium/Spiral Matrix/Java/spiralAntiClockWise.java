class Solution {
    public List<Integer> spiralOrder(int[][] matrix) {
        int r=matrix.length, c=matrix[0].length;
        int left = 0, right = c-1, up = 0, down = r-1;
        List<Integer> res = new ArrayList<>();

        while (res.size() < r*c) {
            for (int i=up; i<=down; i++) res.add(matrix[i][left]);
            for (int i=left+1; i<=right; i++) res.add(matrix[down][i]);
            if (left != right)
                for (int i=down-1; i>=up; i--) res.add(matrix[i][right]);
            
            if (up != down)
                for (int i=right-1; i>left; i--) res.add(matrix[up][i]);

            up++;
            down--;
            left++;
            right--;
        }

        return res;
    }
}
