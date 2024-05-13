#include <stdio.h>
#include <stdlib.h>

int main(int argc, char **argv)
{
    const int max_in = 5;
    double num, avg, sum = 0.0;


    for (int i = 0; i < max_in; i++)
    {
        printf("%d enter a number : ", i);
        scanf("%lf", &num);
        if (num < 0.0)
            goto jump;
        sum += num;
    }

jump:
    printf("number is neg\n");
    return 0;   
}
