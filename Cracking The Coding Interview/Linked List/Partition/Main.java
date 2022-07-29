class Node {
    int val;
    Node next = null;
    Node(int val) { this.val = val; }
}

public class Main {
    
    public static void main(String[] args) {
        Node head = new Node(3);
        head.next = new Node(5);
        head.next.next = new Node(8);
        head.next.next.next = new Node(5);
        head.next.next.next.next = new Node(10);
        head.next.next.next.next.next = new Node(2);
        head.next.next.next.next.next.next = new Node(1);
        Node res = partition(head, 5);
        while(res != null){
            System.out.print(res.val + "->" );
            res = res.next;
        }
        System.out.println();
    }

    public static Node partition(Node head, int x) {
        Node p1 = null;
        Node p2 = null;
        Node pt1 = null;
        Node pt2 = null;
        while(head != null) {
            if (head.val < x){
                if (p1 == null) {
                    p1 = new Node(head.val);
                    pt1 = p1;
                }else {
                    pt1.next = new Node(head.val);
                    pt1 = pt1.next;
                }
            }else {
                if (p2 == null) {
                    p2 = new Node(head.val);
                    pt2 = p2;
                }else {
                    pt2.next = new Node(head.val);
                    pt2 = pt2.next;
                }
            }
            head = head.next;
        }
        pt1.next = p2;
        return p1;
    }
}
