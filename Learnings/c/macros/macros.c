#include <stdio.h>
#include <stdlib.h>

#define PRNT(a, b) \
    printf ("1 = %d\n", a); \
    printf("2= %d\n", b);


#define PRNT_CIR_AR(a)\
    printf ("Area of cirle is %lf", a);

#define PI 3.14
#define CIRCLE_AREA(x) ((PI) * (x) * (x))

#define UPTO(i, n) for ((i) = 0; (i) < (n); (i)++) 

int main (int argc, char **argv) 
{
    int i = 0;
    int c = 5;
    
    PRNT(2, 4);
    PRNT_CIR_AR(CIRCLE_AREA(c + 2));
    
    UPTO(i, 10) {
        printf("%d\n", i);
    }
    
    return 0;
}
