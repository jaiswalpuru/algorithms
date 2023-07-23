class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        HashMap<String, List<String>> map = new HashMap<>();
        List<List<String>> res = new ArrayList<>();
        for (String str : strs) {
            char[] c = str.toCharArray();
            Arrays.sort(c);
            String sb = "";
            for (int i=0; i<c.length; i++) sb += c[i];
            map.computeIfAbsent(sb, k -> new ArrayList<>()).add(str);
        }
        for (String s : map.keySet()) res.add(map.get(s));
        return res;
    }
}
