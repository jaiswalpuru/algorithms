import java.util.HashSet;

class Node {
    Node next = null;
    int val;
    Node(int val) { this.val = val; }
}

public class Main {
    
    public static void main(String[] args) {
        Node n = new Node(1);
        n.next = new Node(2);
        n.next.next = new Node(3);
        n.next.next.next = new Node(1);
        removeDups(n);
        while (n != null) {
            System.out.print(n.val + "->");
            n = n.next;
        }
        System.out.println();
    }

    public static void removeDups(Node n) {
        HashSet<Integer> visited = new HashSet<Integer>();
        Node prev = null;
        while(n != null){
            if (visited.contains(n.val)) {
                prev.next = n.next;
            }else {
                visited.add(n.val);
                prev = n;
            }
            n = n.next;
        }
    }
}
