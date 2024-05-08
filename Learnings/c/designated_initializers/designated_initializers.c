#include <stdio.h>
#include <stdlib.h>

struct Point { int x; int y; int z; };

int main(int argc, char **argv)
{

    //array
    int a[] = {[0 ... 9] = 1, [10 ... 99] = 2, [100] = 3};
    for (int i = 0; i < 101; i++) 
    {
        printf("%d ", a[i]);
    }
    printf("\n");


    int num[] = {1, 2, 3, [10] = 80, 15, [70] = 50, [42] = 400};
    int n = sizeof(num)/sizeof(num[0]);
    printf("size = %d", n);
    for (int i = 0; i < n; i++)
    {
        printf("%d ", num[i]);
    }
    printf("\n");



    // method 1
    struct Point P = {.y =2, .x=1, .z=100}; // or can be written as {y:2, x:1, z:100}
    printf("%d %d %d\n", P.x, P.y, P.z);
    // or

    struct Point P2 = {.x = 100}; // in this case only x will be initialized with 100 rest will be 0
    printf("%d %d %d\n", P2.x, P2.y, P2.z);
    struct Point P3[5] = {[2].y = 100, [2].x = 6, [0].y = 100};
    for (int i = 0; i < 5; i++)
        printf("%d %d %d\n", P3[i].x, P3[i].y, P3[i].z);
    return 0;
}
