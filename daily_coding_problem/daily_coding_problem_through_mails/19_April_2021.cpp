// An XOR linked list is a more memory efficient doubly linked list.
// Instead of each node holding next and prev fields, it holds a field named both,
// which is an XOR of the next node and the previous node. Implement an XOR linked list;
// it has an add(element) which adds the element to the end, and a get(index) which returns the node at index.

// If using a language that has no pointers (such as Python),
// you can assume you have access to get_pointer and dereference_pointer
// functions that converts between nodes and memory addresses.


#include <iostream>

class XORList{ 
    public:
    int data;
    XORList* both;
};

XORList* XOR(XORList *a, XORList* b){
    return reinterpret_cast<XORList *>(reinterpret_cast<uintptr_t>(a)^reinterpret_cast<uintptr_t>(b));
}

void printList(XORList *head){
    XORList *prev = NULL;
    XORList *curr = head;
    XORList *next;
    std :: cout << "XORlist data : ";
    while (curr != NULL) {
        std :: cout << curr->data << " ";
        next = XOR(prev, curr->both);
        prev = curr;
        curr = next;
    }
    std :: cout << "\n";
}

//get -> return the value at specific index
XORList* get(XORList *head, int index){
    XORList* curr = head;
    XORList* prev = NULL;
    XORList* next;
    if ( head == NULL ) {
        std :: cout<<"xor list is empty\n";
        return NULL;
    }else {
        while ( curr != NULL && index!=0 ) {
            next = XOR(prev, curr->both);
            prev = curr;
            curr = next;
            index--;
        }
        if (index == 0) {
            return prev;
        }else{
            std :: cout<<"index out of bounds\n";
            return NULL;
        }
    }
}

//insert -> appends an element to the end of the XORList
void add(XORList **head, int data){
    XORList* newNode = new XORList();
    newNode->data = data;
    if (*head == NULL) {
        newNode->both = *head;
        *head = newNode;
    } else {
        XORList *curr = *head;
        XORList *prev = NULL;
        XORList *next;
        while(curr != NULL){
            next = XOR(prev,curr->both);
            prev = curr;
            curr = next;
        }
        prev->both = XOR(newNode,prev->both);
        newNode->both = XOR(prev,NULL);
    }
}

int main(){

    XORList *head = NULL;
    add(&head,10);
    add(&head,20);
    add(&head,30);
    add(&head,40);
    printList(head);
    
    XORList *val = NULL;
    val = get(head,4);
    if (val != NULL){
        std :: cout << "Value : "<<val->data<<"\n";
    }

    return 0;
}