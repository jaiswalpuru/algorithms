class Node {
    int data;
    Node next = null;
    Node(int val) { this.data = val; }
}

public class Main {
    
    public static void main(String[] args) {
        Node n1 = new Node(7);
        n1.next = new Node(1);
        n1.next.next = new Node(6);
        Node n2 = new Node(5);
        n2.next = new Node(9);
        n2.next.next = new Node(2);
        Node res = sumList(n1, n2);
        while (res != null) {
            System.out.print(res.data + " ");
            res =res.next;
        }
        System.out.println();
    }

    public static Node sumList(Node n1, Node n2) {
        if (n1 == null) return n2;
        if (n2 == null) return n1;
        Node res = new Node(-1);
        Node temp = res;
        int carry=0;
        while (n1!=null && n2!=null){
            int sum = n1.data + n2.data + carry;
            carry = sum/10;
            temp.next = new Node(sum%10);
            temp = temp.next;
            n1=n1.next;
            n2=n2.next;
        }
        while(n1!=null){
            int sum = n1.data + carry;
            carry = sum/10;
            temp.next = new Node(sum%10);
            temp = temp.next;
            n1=n1.next;
        }
        while(n2!=null){
            int sum = n2.data + carry;
            carry = sum/10;
            temp.next = new Node(sum%10);
            temp = temp.next;
            n2 = n2.next;
        }
        while (carry!=0){
            temp.next = new Node(carry%10);
            temp = temp.next;
            carry=carry/10;
        }
        return res.next;
    }
}
