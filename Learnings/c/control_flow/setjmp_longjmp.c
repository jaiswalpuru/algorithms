#include <stdio.h>
#include <stdlib.h>
#include <setjmp.h>

jmp_buf buf;

void my_function()
{
    printf("in my function\n");    
    longjmp(buf, 1);

    /* Not reached */
    printf("you will never see this line printed\n");
}

int main(int argc, char **argv)
{
    if (setjmp(buf)) printf("back in main\n");
    else {
        printf("first time through my function\n");
        my_function();
    }
   // int i = setjmp(buf);
   // printf("i = %d\n", i);
    //if (i != 0) exit (0);
    //longjmp(buf, 42);
    //printf ("this line will not be printed\n infinte loop\n"); // if we remove the if condition on line 12 then it will infinite loop
    return 0;   
}
