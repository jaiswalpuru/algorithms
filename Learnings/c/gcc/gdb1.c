#include <stdio.h>

int main(int argc, char **argv) 
{
    const int arr[5] = {1, 2, 3, 4, 5};
    int i = 0, sum = 0;

    for (i = 0; i >= 0; ++i)
        sum += arr[i];

    printf("Sum of arr : %d\n", sum);
    return 0;
}
