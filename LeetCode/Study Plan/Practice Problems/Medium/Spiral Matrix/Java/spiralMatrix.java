class Solution {
    public List<Integer> spiralOrder(int[][] matrix) {
        int rowSize = matrix.length;
        int colSize = matrix[0].length;
        int left = 0;
        int right = colSize-1;
        int up = 0;
        int down = rowSize-1;
        List<Integer> spiralArr = new ArrayList<>();

        while(spiralArr.size() < rowSize * colSize) {
            for (int j=left; j<=right; j++) spiralArr.add(matrix[up][j]);

            for (int i=up+1; i<=down; i++) spiralArr.add(matrix[i][right]);

            if (up != down) 
                for (int j=right-1; j>=left; j--) spiralArr.add(matrix[down][j]);

            if (left != right)
                for (int i=down-1; i>up; i--) spiralArr.add(matrix[i][left]);
            
            left++;
            right--;
            up++;
            down--;
        }

        return spiralArr;
    }
}
