#include <stdio.h>
#include <stdlib.h>
#include <limits.h>

typedef struct Stack {
    int top;
    int* arr;
    size_t cap;
} Stack;

Stack* init(size_t cap) {
    Stack *s = (Stack *) malloc(sizeof(Stack));
    s->top = -1;
    s->cap = cap;
    s->arr = (int *) malloc(s->cap*sizeof(int));
    return s;
}

int isFull(Stack *s) { return s->top == s->cap-1; }

int isEmpty(Stack *s) { return s->top == -1; }

void push(Stack *s, int val) {
    if (isFull(s)) {
        printf("Stack full\n");
        return;
    }
    s->arr[++s->top] = val;
    printf("%d inserted\n", val);
}

int pop(Stack *s) {
    if (isEmpty(s)) return INT_MIN;
    return s->arr[s->top--];
}

int peek(Stack *s) {
    if (isEmpty(s)) {
        printf("Stack empty");
        return INT_MIN;
    }
    return s->arr[s->top];
}

int main() {
    Stack *s = init(100);
    push(s, 5);
    push(s, 10);
    push(s, 30);
    printf("top of stack : %d\n", peek(s));
    return 0;
}

