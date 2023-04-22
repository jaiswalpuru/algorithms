/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    
    int size;
    ListNode head;
    HashMap<Integer, Integer> indexValue;
    Random rand;

    public Solution(ListNode head) {
        this.head = head;
        rand = new Random();
        update(head);
    }

    public void update(ListNode head) {
        indexValue = new HashMap<>();
        size = 0;

        while(head != null) {
            indexValue.put(size, head.val);
            size++;
            head = head.next;
        }
    }
    
    public int getRandom() {
        return indexValue.get(rand.nextInt(size));
    }
}

/**
 * Your Solution object will be instantiated and called as such:
 * Solution obj = new Solution(head);
 * int param_1 = obj.getRandom();
 */
