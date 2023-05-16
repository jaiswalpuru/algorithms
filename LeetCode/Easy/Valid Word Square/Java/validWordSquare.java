class Solution {
    public boolean validWordSquare(List<String> words) {
        int maxLen = 0;
        for (String word : words) maxLen = Math.max(maxLen, word.length());
        char[][] charArr = new char[Math.max(words.size(), maxLen)][maxLen];
        int k = 0;
        for (int i=0; i<words.size(); i++) {
            String word = words.get(i);
            for (int j=0; j<word.length(); j++)
                charArr[k][j] = word.charAt(j);
            k++;
        }
        List<String> col = new ArrayList<>();
        StringBuilder sb;
        for (int j=0; j<maxLen; j++) {
            sb = new StringBuilder();
            for (int i=0; i<Math.max(words.size(), maxLen); i++) {
                if (charArr[i][j] != 0) {
                    sb.append(charArr[i][j]);
                }
            }
            col.add(sb.toString());
        }

        for (int i=0; i<words.size(); i++)
            if (!col.get(i).equals(words.get(i))){
                return false;
            }
        
        return true;
    }
}
