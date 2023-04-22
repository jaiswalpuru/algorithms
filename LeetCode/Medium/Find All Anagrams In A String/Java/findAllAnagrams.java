class Solution {
    public List<Integer> findAnagrams(String s, String p) {
        int sLen = s.length();
        int pLen = p.length();
        if (pLen > sLen) return new ArrayList<Integer>();

        List<Integer> index = new ArrayList<>();
        Map<Character, Integer> pCharCnt = new HashMap<>();
        Map<Character, Integer> sCharCnt = new HashMap<>();

        for (Character c : p.toCharArray())
            pCharCnt.put(c, pCharCnt.getOrDefault(c, 0)+1);

        for (int i=0; i<sLen; i++) {
            char ch = s.charAt(i);
            sCharCnt.put(ch, sCharCnt.getOrDefault(ch, 0)+1);
            if (i >= pLen) {
                sCharCnt.put(s.charAt(i-pLen), sCharCnt.get(s.charAt(i-pLen))-1);
                if (sCharCnt.get(s.charAt(i-pLen)) == 0) sCharCnt.remove(s.charAt(i-pLen));
            }
            if (pCharCnt.equals(sCharCnt))
                index.add(i-pLen+1);
        }
        return index;
    }
}
