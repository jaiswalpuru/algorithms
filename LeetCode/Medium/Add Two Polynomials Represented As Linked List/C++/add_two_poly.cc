/**
 * Definition for polynomial singly-linked list.
 * struct PolyNode {
 *     int coefficient, power;
 *     PolyNode *next;
 *     PolyNode(): coefficient(0), power(0), next(nullptr) {};
 *     PolyNode(int x, int y): coefficient(x), power(y), next(nullptr) {};
 *     PolyNode(int x, int y, PolyNode* next): coefficient(x), power(y), next(next) {};
 * };
 */

class Solution {
public:
    PolyNode* addPoly(PolyNode* poly1, PolyNode* poly2) {
        PolyNode* res = new PolyNode(-1, -1);
        PolyNode* temp = res;
        int sum = 0;
        while(poly1 != NULL && poly2 != NULL) {
            if (poly1->power == poly2->power) {
                sum = poly1->coefficient + poly2->coefficient;
                if (sum != 0) {
                    temp->next = new PolyNode(sum, poly1->power);
                    temp = temp->next;
                }
                poly1 = poly1->next;
                poly2 = poly2->next;
            } else if (poly1->power > poly2->power) {
                sum = poly1->coefficient;
                if (sum != 0) {
                    temp->next = new PolyNode(sum, poly1->power);
                    temp = temp->next;
                }
                poly1 = poly1->next;
            } else {
                sum = poly2->coefficient;
                if (sum != 0) {
                    temp->next = new PolyNode(sum, poly2->power);
                    temp = temp->next;
                }
                poly2 = poly2->next;
            }
        }

        while(poly1 != NULL) {
            sum = poly1->coefficient;
            temp->next = new PolyNode(sum, poly1->power);
            temp = temp->next;
            poly1 = poly1->next;
        }

        while(poly2 != NULL) {
            sum = poly2->coefficient;
            temp->next = new PolyNode(sum, poly2->power);
            temp = temp->next;
            poly2 = poly2->next;
        }

        return res->next;
    }
};
