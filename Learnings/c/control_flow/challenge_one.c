#include <stdio.h>
#include <stdlib.h>

int main(int argc, char **argv)
{
    int val = 5;
    int i = 1, j, k = 0;

outer_loop:
    if (i < val)
    {
        j = i;

    inner_loop:
        if (j < val)
        {
            printf(" ");
            j++;
            goto inner_loop;
        }
        else
        {
        innerL_loop2:
            if (k != (2*i)) 
            {
                if (k == 0 || k == (2*i)-3)
                {
                    printf("*");
                }
                printf(" ");
                k++;
                goto innerL_loop2;
            }
            k = 0;
            printf("\n");
            i++;
            goto outer_loop;
        }
    }
    else 
    {
        i = 0;
loop:
        if (i < (2*val)-1)
        {
            printf("*");
            i++;
            goto loop;
        }
    }

    return 0;
}
