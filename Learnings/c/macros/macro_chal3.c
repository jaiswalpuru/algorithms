#include <stdio.h>
#include <stdlib.h>

#define CUBE(x) \
    ((x) * (x) * (x))
#define SQ(x) \
    ((x) * (x))

int main(int argc, char **argv)
{
    printf("cube of 3 is %d\n", CUBE(3));
    printf("square of 3 is %d\n", SQ(3));
    return 0;
}
