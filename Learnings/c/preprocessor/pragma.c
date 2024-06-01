#include <stdio.h>

#pragma GCC poison printf
#pragma GCC warning "hello"
#pragma GCC error "ok"

int main(int argc, char **argv)
{
   // printf("Hello world!\n"); // compiler will give error
    
    
    return 0;
}
