class Node {
public:
    int val;
    Node *prev;
    Node *next;
    Node(int val) : val(val), prev(nullptr), next(nullptr) {}
};

class MyCircularDeque {
    int size;
    int num_ele = 0;
    Node *head;
    Node *tail;
public:
    MyCircularDeque(int k) {
        size = k;
        num_ele = 0;
        head = nullptr;
        tail = head;
    }
    
    void reset() {
        if (num_ele == 0) {
            head = nullptr;
            tail = head;
        }
    }

    bool insertFront(int value) {
        if (num_ele < size) {
            Node *node = new Node(value);
            if (head == nullptr) {
                head = node;
                tail = head;
                head->next = tail;
                tail->prev = head;
            } else {
                node->next = head;
                head->prev = node;
                head = node;
            }
            head->prev = tail;
            tail->next = head;
            num_ele++;
            return true;
        }
        return false;
    }
    
    bool insertLast(int value) {
        if (num_ele < size) {
            if (head == nullptr) {
                return insertFront(value);
            } else {
                Node *node = new Node(value);
                tail->next = node;
                node->prev = tail;
                tail = node;
                tail->next = head;
            }
            num_ele++;
            return true;
        }
        return false;
    }
    
    bool deleteFront() {
        if (head == nullptr) return false;
        head = head->next;
        head->prev = tail;
        tail->next = head;
        num_ele--;
        reset();
        return true;
    }
    
    bool deleteLast() {
        if (tail == nullptr) return false;
        tail = tail->prev;
        tail->next = head;
        head->prev = tail;
        num_ele--;
        reset();
        return true;
    }
    
    int getFront() {
        return (head == nullptr) ? -1 : head->val;
    }
    
    int getRear() {
        return (tail == nullptr) ? -1 : tail->val;
    }
    
    bool isEmpty() {
        return num_ele == 0;
    }
    
    bool isFull() {
        return num_ele == size;
    }
};

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * MyCircularDeque* obj = new MyCircularDeque(k);
 * bool param_1 = obj->insertFront(value);
 * bool param_2 = obj->insertLast(value);
 * bool param_3 = obj->deleteFront();
 * bool param_4 = obj->deleteLast();
 * int param_5 = obj->getFront();
 * int param_6 = obj->getRear();
 * bool param_7 = obj->isEmpty();
 * bool param_8 = obj->isFull();
 */
