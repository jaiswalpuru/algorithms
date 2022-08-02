class Node {
    int data;
    Node next = null;
    Node(int data) { this.data = data; }
}

/*
 * One approach to this problem can be to push all the node (not the value) in a hashmap
 * then traverse the second list and see which ever node is present hashmap return that
 *
 * Another approach will be to first get the length and tail of both the list then first 
 * compare if tail doesn't match then there is no intersection else 
 * n = len(list1) m = len(list2)
 * if n>m move the list1 by n-m pointer ahead
 * if m>n move the list2 by m-n pointer ahead 
 * from there move the node one at a time in both list and comparing.
 *
*/

public class Main {
    
    public static void main(String[] args) {
        Node n1 = new Node(3);
        n1.next = new Node(1);
        n1.next.next = new Node(5);
        n1.next.next.next = new Node(9);
        n1.next.next.next.next = new Node(7);
        n1.next.next.next.next.next = new Node(2);
        n1.next.next.next.next.next.next = new Node(1);
        Node n2 = new Node(4);
        n2.next = new Node(6);
        n2.next.next = n1.next.next.next;
        intersection(n1, n2);
    }

    public static void intersection(Node n1, Node n2) {
        if (getTail(n1) != getTail(n2)) {
            System.out.println("No intersection");
            return;
        }
        int n = getLength(n1);
        int m = getLength(n2);
        if (n>m) {
            int diff = n-m;
            while (diff > 0) {
                n1 = n1.next;
                diff--;
            }
        }

        if (m > n) {
            int diff = m-n;
            while (diff > 0) {
                n2 = n2.next;
                diff--;
            } 
        }

        while (n1 != null && n2 != null) {
            if (n1 == n2) {
                System.out.println("Intersection at node : " + n1.data);
                break;
            }
            n1 = n1.next;
            n2 = n2.next;
        }
    }

    public static Node getTail(Node n) {
        if (n==null) return null;
        while(n.next != null) {
            n = n.next;
        }
        return n;
    }

    public static int getLength(Node n){
        int len = 0;
        while (n != null)  {
            len++;
            n = n.next;
        }
        return len;
    }

}
