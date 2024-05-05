#include <stdio.h>
#include <stdlib.h>

struct my_arr {
    int size;
    int arr[];
};

int main(int argc, char **argv)
{
    int i = 0;
    int size = 0;
    struct my_arr *fla;

    printf("Enter the number of elements for arr : ");
    scanf("%d", &size);

    fla = malloc(sizeof(struct my_arr) + size * sizeof(int));
    fla->arr[0] = 1;
    fla->arr[1] = 2;
    for (; i < size; i++)
    {
        printf("%d ", fla->arr[i]);
    }
    printf("\n");

    return 0;       
}
