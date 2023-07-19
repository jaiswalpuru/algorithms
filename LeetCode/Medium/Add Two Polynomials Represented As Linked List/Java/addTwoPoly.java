/**
 * Definition for polynomial singly-linked list.
 * class PolyNode {
 *     int coefficient, power;
 *     PolyNode next = null;
 
 *     PolyNode() {}
 *     PolyNode(int x, int y) { this.coefficient = x; this.power = y; }
 *     PolyNode(int x, int y, PolyNode next) { this.coefficient = x; this.power = y; this.next = next; }
 * }
 */

class Solution {
    public PolyNode addPoly(PolyNode poly1, PolyNode poly2) {
        PolyNode res = new PolyNode(-1 ,-1);
        PolyNode temp = res;
        int sum = 0;
        while(poly1 != null && poly2 != null) {
            if (poly1.power == poly2.power) {
                sum = poly1.coefficient + poly2.coefficient;
                if (sum != 0) {
                    temp.next = new PolyNode(sum, poly1.power);
                    temp = temp.next;
                }
                poly1 = poly1.next;
                poly2 = poly2.next;
            } else if (poly1.power > poly2.power) {
                sum = poly1.coefficient;
                if (sum != 0) {
                    temp.next = new PolyNode(sum, poly1.power);
                    temp = temp.next;
                }
                poly1 = poly1.next;
            } else {
                sum = poly2.coefficient;
                if (sum != 0) {
                    temp.next = new PolyNode(sum, poly2.power);
                    temp = temp.next;
                }
                poly2 = poly2.next;
            }
        }

        while(poly1 != null) {
            sum = poly1.coefficient;
            temp.next = new PolyNode(sum, poly1.power);
            temp = temp.next;
            poly1 = poly1.next;
        }
        while(poly2 != null) {
            sum = poly2.coefficient;
            temp.next = new PolyNode(sum, poly2.power);
            temp = temp.next;
            poly2 = poly2.next;
        }

        return res.next;
    }
}
