class Solution {
    public boolean canFinish(int numCourses, int[][] prerequisites) {
        HashMap<Integer, List<Integer>> g = new HashMap<>();
        for (int i=0; i<prerequisites.length; i++)
            g.computeIfAbsent(prerequisites[i][1], k->new ArrayList<>()).add(prerequisites[i][0]);
        
        int[] indegree = new int[numCourses];
        for (int[] p : prerequisites) indegree[p[0]]++;
        
        ArrayDeque<Integer> q = new ArrayDeque<>();
        for (int i=0; i<numCourses; i++)
            if (indegree[i]==0) q.add(i);
            
        while(!q.isEmpty()) {
            int curr = q.pop();
            if (g.containsKey(curr)) {
                for (Integer v : g.get(curr)) {
                    indegree[v]--;
                    if (indegree[v] == 0) q.add(v);
                }
            }
        }
    
        for (int i=0; i<numCourses; i++) {
            if (indegree[i] > 0) return false;
        }

        return true;
    }
}
