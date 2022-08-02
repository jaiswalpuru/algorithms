class Node {
    int data;
    Node next = null;
    Node(int data) { this.data = data; }
}

public class Main {
    
    public static void main(String[] args) {
        Node n = new Node(1);
        n.next = new Node(2);
        n.next.next = new Node(3);
        n.next.next.next = new Node(4);
        n.next.next.next.next = n.next;
        loopDetection(n);
    }

    /*
     * we can either use hashmap and check if a node is alreday visited, then there is loop
     * or
     * use two pointer approach
    */
    public static void loopDetection(Node n) {
        if (n==null) {
            System.out.println("No loop");
            return;
        }
        Node p1 = n;
        Node p2 = n;
        while(p2 != null && p2.next != null){
            p1 = p1.next;
            p2 = p2.next.next;
            if (p1==p2) {
                System.out.println("loop detected : " + p2.next.data);
                return;
            }
        }
        System.out.println("No loop");
    }
}
