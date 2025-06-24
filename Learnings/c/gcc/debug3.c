#include <stdio.h>
#include <stdlib.h>

#define DEBUG(level, fmt, ...) \
    if (level >= 2) \
        fprintf(stderr, fmt, __VA_ARGS__) 

int main(int argc, char **argv) 
{
    int val = 0;


    return 0;
}
