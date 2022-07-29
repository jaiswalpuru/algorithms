class Node {
    Node next = null;
    int val;
    Node(int val) { this.val = val; }
}

public class Main {
    
    public static void main(String[] args) {
        Node head = new Node(1);
        head.next = new Node(2);
        head.next.next = new Node(3);
        head.next.next.next = new Node(4);
        deleteNode(head.next);
    }

    public static void deleteNode(Node head) {
        if (head == null || head.next == null) {
            return;
        }
        Node temp = head.next;
        head.val = temp.val;
        head.next = temp.next;
        return;
    }
}
