#include <stdio.h>
#include <stdlib.h>

typedef struct Node {
    int val;
    struct Node* next;
}Node;

void print_list(Node* list) {
    while(list != NULL) {
        printf("%d->", list->val);
        list = list->next;
    }
    printf("\n");
}

void insert_front(Node** list, int val) {
    Node *new_node = (Node*) malloc(sizeof(Node));
    new_node->val = val;
    new_node->next = *list;
    *list = new_node;
}

void insert_back(Node **list, int val) {
    Node* new_node = (Node*) malloc(sizeof(Node));
    new_node->val = val;

    if (*list == NULL) {
        *list = new_node;
        return;
    }

    Node *temp = *list;
    while(temp->next != NULL) temp = temp->next;
    temp->next = new_node;
}

void delete_node(Node **head, int val) {
    Node *temp = *head, *prev;
    if (temp != NULL && temp->val == val) {
        *head = temp->next;
        free(temp);
        return;
    }
    while (temp != NULL && temp->val != val) {
        prev = temp;
        temp = temp->next;
    }

    if (temp == NULL) return;
    prev->next = temp->next;
    free(temp);
}

void sort_list(Node **list) {
    if (*list == NULL) return;
    Node* curr = *list, *index = NULL;
    while(curr != NULL) {
        index = curr->next;
        while(index != NULL) {
            if (curr->val > index->val) {
                int temp = curr->val;
                curr->val = index->val;
                index->val = temp;
            }
            index = index->next;
        }
        curr = curr->next;
    }
}

int search_node(Node **head, int val) {
    Node* temp = *head;
    while(temp != NULL) {
        if (temp->val == val) return val;
        temp = temp->next;
    }
    return -1; // indicates element not found    
}

int main() {
    Node* head = (Node*) malloc(sizeof(Node));
    head->val = 4;
    head->next = (Node*) malloc(sizeof(Node));
    head->next->val = 3;
    head->next->next = (Node*) malloc(sizeof(Node));
    head->next->next->val = 2;
    head->next->next->next = (Node*) malloc(sizeof(Node));
    head->next->next->next->val = 1; 
    print_list(head);
    sort_list(&head);
    insert_front(&head, 10);
    print_list(head);
    insert_back(&head, 34);
    print_list(head);
    delete_node(&head, 10);
    print_list(head);
    delete_node(&head, 4);
    print_list(head);
    sort_list(&head);
    print_list(head);
}
