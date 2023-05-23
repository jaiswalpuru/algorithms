class P<T> {
    T a, b;
    public P(T a, T b) {
        this.a = a;
        this.b = b;
    }
}

class Solution {
    public int[] topKFrequent(int[] nums, int k) {
        PriorityQueue<P<Integer>> pq = new PriorityQueue<>((a, b) -> { return b.b-a.b; });
        Map<Integer, Integer> cntMap = new HashMap<>();
        for (int i=0; i<nums.length; i++)
            cntMap.put(nums[i], cntMap.getOrDefault(nums[i], 0)+1);
        
        for (Map.Entry<Integer, Integer> m : cntMap.entrySet())
            pq.offer(new P<Integer>(m.getKey(), m.getValue()));

        int[] res = new int[k];
        int i=0;
        while(k-- > 0) res[i++] = pq.poll().a;
        return res;
    }
}


// ******* using quickselect *******
class Solution {

    Map<Integer, Integer> cntMap;
    int[] unique;

    public int[] topKFrequent(int[] nums, int k) {
        cntMap = new HashMap<>();
        for (int i=0; i<nums.length; i++)
            cntMap.put(nums[i], cntMap.getOrDefault(nums[i], 0)+1);

        int i = 0;
        int n = cntMap.size();
        unique = new int[n];
        for (int num : cntMap.keySet()) {
            unique[i++] = num;
        }

        quickSelect(0, n-1, n-k);
        return Arrays.copyOfRange(unique, n-k, n);
    }

    private void quickSelect(int left, int right, int kSmallest) {
        if (left == right) return;

        Random r = new Random();
        int pivotIndex = left+r.nextInt(right-left);
        pivotIndex = partition(left, right, pivotIndex);
        if (kSmallest == pivotIndex) return;
        else if (kSmallest < pivotIndex) quickSelect(left, pivotIndex-1, kSmallest);
        else quickSelect(pivotIndex+1, right, kSmallest);
    }

    private int partition(int left, int right, int pivotIndex) {
        //move pivot to end
        int pivotFreq = cntMap.get(unique[pivotIndex]);
        swap(pivotIndex, right);
        int storeIndex = left;

        //move all the less frequent elements to the left
        for (int i=left; i<=right; i++) {
            if (cntMap.get(unique[i]) < pivotFreq) {
                swap(storeIndex, i);
                storeIndex++;
            }
        }

        swap(storeIndex, right);
        return storeIndex;
    }

    private void swap(int i, int j) {
        int temp = unique[i];
        unique[i] = unique[j];
        unique[j] = temp;
    }
}
