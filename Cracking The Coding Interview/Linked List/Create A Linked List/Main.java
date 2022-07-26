public class Main {

    public static void main(String[] args) { }

    public static Node deleteNode(Node head, int d) {
        if (head == null) return null;
        Node n = head;
        if (n.val == d) { return n.next; }
        while (n.next != null) {
            if (n.next.val == d) {
                n.next = n.next.next;
                return head;
            }
            n = n.next;
        }
        return head;
    }
}
