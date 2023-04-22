class Solution {
    class Project implements Comparable<Project> {
        int profit;
        int capital;

        public Project(int profit, int capital) {
            this.profit = profit;
            this.capital = capital;
        }

        @Override
        public int compareTo(Project p) { return capital-p.capital; }
    }

    public int findMaximizedCapital(int k, int w, int[] profits, int[] capital) {
        int size = profits.length;
        Project[] projects = new Project[size];

        for (int i=0; i<size; i++) {
            projects[i] = new Project(profits[i], capital[i]);
        }    
        
        Arrays.sort(projects);

        PriorityQueue<Integer> maxHeap = new PriorityQueue<>(size, Collections.reverseOrder());

        int ptr = 0;
        for (int i=0; i<k; i++) {
            while(ptr < size && projects[ptr].capital <= w) {
                maxHeap.offer(projects[ptr].profit);
                ptr++;
            }
            if (maxHeap.isEmpty()) break;
            w += maxHeap.poll();
        }

        return w;
    }
}
