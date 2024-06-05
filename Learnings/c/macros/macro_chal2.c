#include <stdio.h>
#include <stdlib.h>

#define SUM(a, b) \
    (a) + (b)

int main(int argc, char **argv)
{
    printf("sum of 2 and 3 = %d\n", SUM(2, 3));
    return 0;
}
