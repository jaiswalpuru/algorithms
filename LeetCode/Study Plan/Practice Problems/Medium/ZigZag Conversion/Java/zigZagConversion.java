class Solution {
    public String convert(String s, int numRows) {
        if (numRows == 1) return s;

        int size = s.length();
        int sections = (int)Math.ceil(size/(2*numRows-2.0));
        int numCols = sections * (numRows-1);

        char[][] table = new char[numRows][numCols];
        for (char[] rows: table) {
            Arrays.fill(rows, ' ');
        }

        int currRow = 0, currCol = 0;
        int index = 0;

        while(index < size) {
            //down
            while(currRow < numRows && index < size) {
                table[currRow][currCol] = s.charAt(index);
                index++;
                currRow++;
            }

            currRow -= 2;
            currCol++;

            while(currRow > 0 && index < size && currCol < numCols) {
                table[currRow][currCol] = s.charAt(index);
                index++;
                currCol++;
                currRow--;
            }
        }

        StringBuilder res = new StringBuilder();
        for (int i = 0; i < numRows; i++) {
            for (int j = 0; j < numCols; j++) {
                if (table[i][j] != ' ') {
                    res.append(table[i][j]);
                }
            }
        }
        return res.toString();
    }
}
