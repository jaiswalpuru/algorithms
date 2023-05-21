class Solution {
    public boolean closeStrings(String word1, String word2) {
        if (word1.length() != word2.length()) return false;
        Map<Character, Integer> word1Char = new HashMap<>();
        Map<Character, Integer> word2Char = new HashMap<>();
        for (int i=0; i<word1.length(); i++) {
            word1Char.put(word1.charAt(i), word1Char.getOrDefault(word1.charAt(i), 0) + 1);
            word2Char.put(word2.charAt(i), word2Char.getOrDefault(word2.charAt(i), 0) + 1);
        }

        Map<Integer, Integer> cntOne = new HashMap<>();
        Map<Integer, Integer> cntTwo = new HashMap<>();

        for (Character c : word1Char.keySet()) {
            if (!word2Char.containsKey(c)) return false;
            int count = word1Char.get(c);
            cntOne.put(count, cntOne.getOrDefault(count, 0) + 1);
        }

        for (Character c : word2Char.keySet()) {
            int count = word2Char.get(c);
            cntTwo.put(count, cntTwo.getOrDefault(count, 0) + 1);
        }

        for (Integer k : cntOne.keySet()) {
            if (cntOne.get(k) != cntTwo.get(k)) return false;
        }

        return true;
    }
}
