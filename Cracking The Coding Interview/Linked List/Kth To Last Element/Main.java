class Node {
    Node next = null;
    int val;
    Node(int val) { this.val=val; }
}

public class Main{
    
    public static void main(String[] args) {
        Node n = new Node(1);
        n.next = new Node(2);
        n.next.next = new Node(3);
        n.next.next.next = new Node(4);
        System.out.println(kToLastNode(n, 2));
        System.out.println(kToLastNodeIterative(n, 2).val);
    }

    public static int kToLastNode(Node n, int k) {
        if (n==null) return 0;
        int ind = kToLastNode(n.next, k)+1;
        if (ind==k){ System.out.println("Kth Node is :" + n.val); }
        return ind;
    }

    public static Node kToLastNodeIterative(Node n, int k) {
        Node p1 = n;
        Node p2 = n;
        for (int i=0; i<k; i++) {
            if (p1 == null) return null;
            p1 = p1.next;
        }
        while(p1 != null){
            p1 = p1.next;
            p2 = p2.next;
        }
        return p2;
    }
    
}
