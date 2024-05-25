#include <stdio.h>
#include <stdlib.h>

union mixed {
    char c;
    float f;
    int i;
};

union member {
    int x;
    double y;
};

void foo(union member n) {
    union member m = n;
}

int main(int argc, char **argv)
{
    union member mem;
    union mixed x;
    x.c = 'j'; // only one member can be stored in an union
    printf("%c\n", x.c);
    mem.x = 100;
    printf("member mem.x = %d\n", mem.x);
    printf("member mem.y = %f\n", mem.y);
    // only one is accessed at any given time
    mem.y = 100;
    printf("member mem.x = %d\n", mem.x);
    printf("member mem.y = %f\n", mem.y);


    // can also be initialized like below
    union member mem2 = {15}; // union can be initialized with the type of first member
                              // if done like union member mem2 ={14.54}; this will throw compiler warning

    // or union member mem2 = {.x = 15};
    //
    return 0;
}
