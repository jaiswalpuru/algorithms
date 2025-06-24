#include <stdio.h>

void new_func1(void) {
    printf("inside new func 1\n");
    int i = 0;

    for (; i < 0xffffffee; i++);
    return;
}
