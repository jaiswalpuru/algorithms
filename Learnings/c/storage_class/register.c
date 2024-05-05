#include <stdio.h>
#include <stdlib.h>

int main(int argc, char **argv)
{
    register int x;
    for (x = 1; x < 15; x++)
    {
        printf("%d ", x);
    }
    printf("\n");
    return 0;
}
