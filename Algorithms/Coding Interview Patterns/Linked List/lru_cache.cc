#include <iostream>
#include <unordered_map>
#include "doubly_ll.h"

using namespace std;

class LRU {
    unordered_map<int, Node*> hash_map;
    int cap;
    Node* head;
    Node* tail;

public:
    LRU(int cap) {
        this->cap = cap;
        this->head = new Node;
        this->tail = new Node;
        this->head->key = -1;
        this->head->val = -1;
        this->tail->key = -1;
        this->tail->val = -1;
        this->head->next = this->tail;
        this->tail->prev = this->head;
    }

    int get(int key) {
        if (this->hash_map.find(key) == this->hash_map.end()) return -1;
        
        remove(this->hash_map[key]);
        add_to_tail(this->hash_map[key]);
        return this->hash_map[key]->val;
    }

    void put(int key, int val) {
        if (this->hash_map.find(key) != this->hash_map.end()) remove(this->hash_map[key]);
        Node* node = new Node;
        node->key = key;
        node->val = val;
        this->hash_map[key] = node;

        if (this->hash_map.size() > this->cap) {
            this->hash_map.erase(this->head->next->key);
            remove(this->head->next);
        }
        add_to_tail(node);
    } 

    void add_to_tail(Node* node) {
        Node* prev = this->tail->prev;
        node->prev = prev;
        node->next = this->tail;
        prev->next = node;
        this->tail->prev = node;
    }

    void remove(Node* node) {
        node->prev->next = node->next;
        node->next->prev = node->prev;
    }
};

int main(int argc, char **argv) {
    LRU lru(3);

    lru.put(2, 250);
    lru.put(4, 300);
    lru.put(3, 200);
    cout << lru.get(2) << "\n";
    cout << lru.get(3) << "\n";
    return 0;
}

