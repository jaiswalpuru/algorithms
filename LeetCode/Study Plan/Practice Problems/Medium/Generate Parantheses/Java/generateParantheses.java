class Solution {
    public List<String> generateParenthesis(int n) {
        List<String> res = new ArrayList<>();
        StringBuilder sb = new StringBuilder();
        recur(res, sb, 0, 0, n);
        return res;
    }

    private void recur(List<String> res, StringBuilder sb, int i, int j, int n) {
        if (i+j == 2*n) {
            res.add(new String(sb.toString()));
            return;
        }
        
        if (i < n) {
            sb.append("(");
            recur(res, sb, i+1, j, n);
            sb.deleteCharAt(sb.length()-1);
        }

        if (i > j) {
            sb.append(")");
            recur(res, sb, i, j+1, n);
            sb.deleteCharAt(sb.length()-1);
        }

    }
}
