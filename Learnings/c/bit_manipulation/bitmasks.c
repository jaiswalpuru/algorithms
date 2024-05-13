#include <stdio.h>
#include <stdlib.h>


int main(int argc, char **argv)
{
    int flag = 15;  // 0000 1111
    int mask = 182; // 1011 0110

    printf("%d\n", (flag | mask));
    printf("%d\n", (flag & ~(mask)));
    printf("%d\n", (flag ^ mask));

    return 0;
}
