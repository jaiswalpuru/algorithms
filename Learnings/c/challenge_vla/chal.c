#include <stdio.h>

int sum1d(size_t, int[*]);

int main(int argc, char **argv)
{
    size_t size = 0;
    printf("Enter the sizeof array : ");
    scanf("%zd", &size);
    int arr[size];
    printf("Enter %zd elements to be stored\n", size);
    int i = 0;
    for (; i < size; i++)
    {
        scanf("%d", &arr[i]);
    }
    printf("Sum of elements : %d", sum1d(size, arr));
    return 0;
}

int sum1d(size_t size, int arr[size])
{
    int sum = 0;
    int i = 0;
    for (; i < size; i++)
    {
        sum += arr[i];    
    }
    return sum;
}
