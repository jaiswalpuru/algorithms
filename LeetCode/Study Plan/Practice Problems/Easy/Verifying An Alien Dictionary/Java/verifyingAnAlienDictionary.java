class Solution {
    HashMap<Character, Integer> orderNum = new HashMap<>();

    public boolean isAlienSorted(String[] words, String order) {
        int k = 0;
        for (Character c : order.toCharArray()) {
            orderNum.put(c, k++);
        }

        int size = words.length;
        for (int i=1; i < size; i++) {
            String w1 = words[i-1];
            String w2 = words[i];
            if (!isSorted(w1, w2))
                return false;
        }
        return true;
    }

    public boolean isSorted(String s1, String s2) {
        int i=0, j=0;
        
        while (i<s1.length() && j<s2.length() && s1.charAt(i) == s2.charAt(j)) {
            i++;
            j++;
        }

        if (((i<s1.length()) && (j == s2.length())) || 
        ((i<s1.length() && j<s2.length()) && (orderNum.get(s1.charAt(i)) > orderNum.get(s2.charAt(j)))))
            return false;

        return true;
    }
}
