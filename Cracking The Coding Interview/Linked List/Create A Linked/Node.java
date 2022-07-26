public class Node {
    Node next = null;
    int val;

    public Node(int d) { val = d; }
    public void addToEnd(int d) {
        Node end = new Node(d);
        Node n = this;
        while (n.next != null) n = n.next;
        n.next = end;
    }
}
