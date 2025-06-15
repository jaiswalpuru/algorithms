typedef struct QueueNode {
    int val;
    QueueNode* next;
    QueueNode(int v) : val(v), next(nullptr) {}
} Node;

class BoundedBlockingQueue {
    Node* head;
    Node* tail;
    condition_variable cv;
    mutex mtx;
    int cap;
    int sz;
public:
    BoundedBlockingQueue(int capacity) {
        head = nullptr;
        tail = nullptr;
        sz = 0;
        cap = capacity;
    }
    
    void enqueue(int element) {
        unique_lock<mutex> lock(mtx);
        cv.wait(lock, [this]{ return sz < cap; });
        Node *new_node = new Node(element);
        if (head == nullptr) {
            head = new_node;
            tail = new_node;
        } else {
            tail->next = new_node;
            tail = new_node;
        }
        sz++;
        cv.notify_all();
    }
    
    int dequeue() {
        unique_lock<mutex> lock(mtx);
        cv.wait(lock, [this]{ return sz > 0; });
        int val = head->val;
        head = head->next;
        sz--;
        if (sz == 0) {
            head = nullptr;
            tail = nullptr;
        }
        cv.notify_all();
        return val;
    }
    
    int size() {
        return sz;
    }
};
