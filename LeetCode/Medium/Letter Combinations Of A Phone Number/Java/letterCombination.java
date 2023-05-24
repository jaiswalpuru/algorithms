class Solution {
    Map<Character, String> numStr = new HashMap<>(){{
        put('2', "abc");
        put('3', "def");
        put('4', "ghi");
        put('5', "jkl");
        put('6', "mno");
        put('7', "pqrs");
        put('8', "tuv");
        put('9', "wxyz");
    }};
    List<String> res;
    public List<String> letterCombinations(String digits) {
        res = new ArrayList<>();
        backtrack(digits, new StringBuilder(), 0);
        return res;
    }

    private void backtrack(String digits, StringBuilder temp, int index) {
        if (index == digits.length()) {
            if (temp.length() > 0)
                res.add(new String(temp.toString()));
            return;
        }
        char[] c = numStr.get(digits.charAt(index)).toCharArray();
        for (int i=0; i<c.length; i++) {
            temp.append(c[i]);
            backtrack(digits, temp, index+1);
            temp.deleteCharAt(temp.length()-1);
        }
    }
}
