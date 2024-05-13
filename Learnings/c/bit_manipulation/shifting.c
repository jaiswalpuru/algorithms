#include <stdio.h>
#include <stdlib.h>

int main(int argc, char **argv)
{
    int w1 = 3;
    int res = 0;


    res = w1 << 1; // multiply by 2;

    printf("updated val  : %d\n", res);

    w1 = 138;
    res = w1 >> 2; // divide by 2^2


    printf("updated val after right shift : %d", res);

    return 0;
}
