#include <stdio.h>

typedef int* i_ptr;

#define peach int
unsigned peach i;
typedef int banana;

//unsigned banana i; // illegal as i is already there


#define i_ptr int *

i_ptr chalk, cheese; //same as int *chalk, cheese
                     //

typedef int* int_ptr;
int_ptr c, d; // same as int *c, *d;;

int main(int argc, char **argv)
{
    i_ptr p; // int *p
    i_ptr a, *b; // same as int *a, int **b;
    i_ptr arr[10]; // int *arr[10]
    return 0;
}
