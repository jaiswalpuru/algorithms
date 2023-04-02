class Solution {
    public List<String> letterCasePermutation(String s) {
        List<String> res = new ArrayList<>();
        backtrack(s, res, "", 0);
        return res;       
    }

    public void backtrack(String s, List<String> res, String temp, int ind) {
        if (ind == s.length()) {
            res.add(new String(temp));
            return;
        }

        char c = s.charAt(ind);
        if (!Character.isDigit(c)) {
            if (Character.isUpperCase(c))
                backtrack(s, res, temp+Character.toLowerCase(c), ind+1);
            else
                backtrack(s, res, temp+Character.toUpperCase(c), ind+1);
        }
        backtrack(s, res, temp+c, ind+1);
    }
}
