#include <stdio.h>
#include <stdlib.h>
#include <setjmp.h>

jmp_buf buf;

void err_function() {
    printf("detected unrecoverable error \n");
    longjmp(buf, 1);
}

int main(int argc, char **argv) {
    
    while(1) {
        if (setjmp(buf)) {
            printf("back in main\n");
            break;
        } else {
            err_function();
        }
    }
    return 0;
}
