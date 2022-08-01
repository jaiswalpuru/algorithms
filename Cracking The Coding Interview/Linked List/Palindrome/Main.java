class Node {
    int data;
    Node next = null;
    Node(int data) { this.data = data; }
}

public class Main {
    
    private static Node temp;
    public static void main(String[] args) {
        Node head = new Node(1);
        head.next = new Node(2);
        head.next.next = new Node(1);
        temp = head;
        System.out.println(isPalindrome(head));
    }

    public static boolean isPalindrome(Node head) {
        if (head != null) {
            if (!isPalindrome(head.next)) {
                return false;
            }
            if (head.data!=temp.data) {
                return false;
            }
            temp = temp.next;
        }
        return true;
    }
}
