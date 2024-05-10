#include <stdio.h>
#include <stdlib.h>


void f1(int n, float * restrict a1, const float * restrict a2)
{
    int i;
    for (i = 0; i < n; i++)
    {
        a1[i] += a2[i];
    }
}

int main(int argc, char **argv)
{
    int n;
    int arr[10] = {[0 ... 9] = 0};
    int * restrict restar = (int *)malloc(10 * sizeof(int));
    int *par = arr; // cannot be qualified as restrict
                    //
                    //

    for (n = 0; n < 10; n++)
    {
        par[n] += 5;
        restar[n] += 5;
        arr[n] *= 2;
        par[n] += 3;
        restar[n] += 3;
    }
    
    // in the above loop compiler can do optimization that, restar[n] += 8, but compiler wont do optimization on par and arr as arr is used in between par modification

    for (n = 0; n < 10; n++)
        printf("%d ", par[n]);

    printf("\n");
    return 0;
}
