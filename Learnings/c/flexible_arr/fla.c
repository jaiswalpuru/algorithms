#include <stdio.h>
#include <stdlib.h>

struct s {
    int size;
    int arr[];
}; // end of struct

int main(int argc, char **argv)
{
    int desired_size = 4;
    struct s *ptr;
    ptr = malloc(sizeof(struct s) + desired_size * sizeof(int));

    return 0;
} 
