class Solution {
    HashMap<String, Boolean> visited;
    public boolean isScramble(String s1, String s2) {
        visited = new HashMap<>();
        return recur(s1, s2);
    }

    public boolean recur(String s1, String s2) {
        int size = s1.length();
        if (s1.equals(s2)) return true;
        
        if (size == 1) return false;

        String memo = s1 + ":" + s2;
        if (visited.containsKey(memo)) return visited.get(memo);

        for (int i=1; i<size; i++) {
            boolean dontSwap = recur(s1.substring(0, i), s2.substring(0, i)) && 
                            recur(s1.substring(i), s2.substring(i));
                        
            if (dontSwap) {
                visited.put(memo, true);
                return true;
            }

            boolean swap = recur(s1.substring(0, i), s2.substring(size-i)) &&
                            recur(s1.substring(i), s2.substring(0, size-i));
            
            if (swap) {
                visited.put(memo, true);
                return true;
            }
        }
        visited.put(memo, false);
        return visited.get(memo);
    }
}
