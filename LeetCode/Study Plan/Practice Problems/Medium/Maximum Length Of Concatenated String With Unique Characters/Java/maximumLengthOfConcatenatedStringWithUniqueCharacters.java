class Solution {
    int res = 0;

    public int maxLength(List<String> arr) {
        backtrack(arr, 0, new ArrayList<>());
        return res;    
    }

    public void backtrack(List<String> arr, int index, List<String> set) {
        int currLength = 0;
        for (int i=index; i<arr.size(); i++) {
            set.add(arr.get(i));
            String joined = String.join("", set);
            if (isUnique(joined)){
                currLength = joined.length();
                backtrack(arr, i+1, set);
            }
            res = Math.max(res, currLength);
            set.remove(set.size()-1);
        }
    }

    public boolean isUnique(String s) {
        HashMap<Character, Integer> hashMap = new HashMap<>();
        for (int i=0; i<s.length(); i++) {
            hashMap.put(s.charAt(i), hashMap.getOrDefault(s.charAt(i), 0)+1);
            if (hashMap.get(s.charAt(i)) > 1)
                return false;
        }
        return true;
    }
}
