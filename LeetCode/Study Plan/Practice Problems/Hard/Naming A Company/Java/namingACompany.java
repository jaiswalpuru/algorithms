class Solution {
    public long distinctNames(String[] ideas) {
        long res = 0;
        HashSet<String>[] groups = new HashSet[26];
        for (int i=0; i<26; i++) {
            groups[i] = new HashSet();
        }

        for (String name : ideas) {
            groups[name.charAt(0)-'a'].add(name.substring(1));
        }

        for (int i=0; i<26; i++) {
            for (int j=i+1; j<26; j++) {
                
                long mutual = 0;
                for (String name : groups[i]) {
                    if (groups[j].contains(name)) mutual++;
                }

                res += 2 * (groups[i].size()-mutual) * (groups[j].size() - mutual);
            }
        }

        return res;
    }
}
