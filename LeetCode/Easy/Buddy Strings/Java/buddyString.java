class Solution {
    public boolean buddyStrings(String s, String goal) {
        int n = s.length();
        int m = goal.length();
        if (n != m) return false;

        if (s.equals(goal)) {
            int[] freq = new int[26];
            for (char c : s.toCharArray()) {
                freq[c-'a']++;
                if (freq[c-'a'] == 2) return true;
            }
            return false;
        }

        int firstIndex = -1;
        int secondIndex = -1;
        for (int i=0; i<s.length(); i++) {
            if (s.charAt(i) != goal.charAt(i)) {
                if (firstIndex == -1) firstIndex = i;
                else if (secondIndex == -1) secondIndex = i;
                else return false;
            }
        }
        if (secondIndex == -1) return false;
        return s.charAt(firstIndex) == goal.charAt(secondIndex) && s.charAt(secondIndex) == goal.charAt(firstIndex);
    }
}
