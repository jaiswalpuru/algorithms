class Solution {
    public String predictPartyVictory(String senate) {
        Queue<Integer> radiant = new LinkedList<>();
        Queue<Integer> dire = new LinkedList<>();
        int n = senate.length();

        for (int i=0; i<n; i++) {
            if (senate.charAt(i) == 'R') radiant.add(i);
            else dire.add(i);
        }

        while (!radiant.isEmpty() && !dire.isEmpty()) {
            int d = dire.poll();
            int r = radiant.poll();

            if (d < r) {
                dire.add(r+n);
            }else {
                radiant.add(d+n);
            }
        }

        return dire.isEmpty() ? "Radiant" : "Dire";
    }
}
