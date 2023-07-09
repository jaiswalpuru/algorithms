class Solution {
    public int maxConsecutiveAnswers(String answerKey, int k) {
        int maxSize = k;
        Map<Character, Integer> map = new HashMap<>();
        for (int i=0; i<k; i++)
            map.put(answerKey.charAt(i), map.getOrDefault(answerKey.charAt(i), 0)+1);
        
        int left = 0;
        for (int right = k; right < answerKey.length(); right++) {
            map.put(answerKey.charAt(right),
                    map.getOrDefault(answerKey.charAt(right), 0)+1);
                
            while(Math.min(map.getOrDefault('T', 0), map.getOrDefault('F', 0)) > k) {
                map.put(answerKey.charAt(left), map.getOrDefault(answerKey.charAt(left), 0)-1);
                left++;
            }
            maxSize = Math.max(maxSize, right-left+1);
        }
        return maxSize;
    }
}
