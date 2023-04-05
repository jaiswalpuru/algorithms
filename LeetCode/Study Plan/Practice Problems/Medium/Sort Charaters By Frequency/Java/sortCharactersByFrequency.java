//------------------------method 1-----------------------------------
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
//------------------------method 1-----------------------------------

//------------------------method 2-----------------------------------
class Solution {
    public String frequencySort(String s) {
        PriorityQueue<Map.Entry<Character, Integer>> pq = new PriorityQueue<>(new Comparator<>(){
            public int compare(Map.Entry<Character, Integer> a, Map.Entry<Character, Integer> b) {
                return b.getValue()-a.getValue();
            }
        });

        HashMap<Character, Integer> cntMap = new HashMap<>();
        for (Character c : s.toCharArray())
            cntMap.put(c, cntMap.getOrDefault(c, 0)+1);

        for (Map.Entry<Character, Integer> m : cntMap.entrySet()) 
            pq.add(m);
        
        StringBuilder sb = new StringBuilder();
        while(!pq.isEmpty()) {
            Map.Entry<Character, Integer> curr = pq.poll();
            int cnt = curr.getValue();
            while(cnt-- > 0) {
                sb.append(curr.getKey());
            }
        }

        return sb.toString();
    }
}
//------------------------method 2-----------------------------------
