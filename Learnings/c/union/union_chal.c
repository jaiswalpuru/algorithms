#include <stdio.h>
#include <stdlib.h>

union student {
    char lg;
    int rg;
    float eg;
};

int main(int argc, char **argv)
{
    union student v1, v2;
    v1.lg = 'B';
    v1.rg = 87;
    v1.eg = 86.50;

    printf("letter : %c\n", v1.lg);
    printf("rounded grade : %d\n", v1.rg);
    printf("exact : %f\n", v1.eg);

    v2.lg = 'B';
    printf("letter v2: %c\n", v2.lg);
    v2.rg = 87;
    printf("rounded v2: %d\n", v2.rg);
    v2.eg = 86.50;
    printf("exact v2 : %f\n", v2.eg);
    return 0;
}

