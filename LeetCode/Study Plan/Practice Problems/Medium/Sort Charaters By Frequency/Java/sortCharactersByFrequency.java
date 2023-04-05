class Solution {
    public String frequencySort(String s) {
        HashMap<Character, Integer> cntMap = new HashMap<>();
        Map<Integer, List<Character>> revMap = new TreeMap<>(Collections.reverseOrder());
        
        for (char c : s.toCharArray()) 
            cntMap.put(c, cntMap.getOrDefault(c, 0)+1);

        for (Map.Entry<Character, Integer> m : cntMap.entrySet())
            revMap.computeIfAbsent(m.getValue(), k->new ArrayList<>()).add(m.getKey());

        StringBuilder sb = new StringBuilder();
        for (Integer val : revMap.keySet()) {
            List<Character> res = revMap.get(val);
            for (Character c : res) {
                int cnt = val;
                while(cnt-- > 0) sb.append(c);
            }
        }

        return sb.toString();
    }
}
