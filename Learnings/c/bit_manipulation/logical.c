#include <stdlib.h>
#include <stdio.h>

int main(int argc, char **argv)
{
    short int w1 = 25;
    short int w2 = 77;
    short int w3 = 0;

    w3 = w1 & w2; // 9
    printf("%d \n", w3);

    printf("%d\n", w1 | w2);


    // to swap
    short int temp = 0;
    temp = w1;
    w1 = w2;
    w2 = temp;

    // rather doing the above we can use xor operator to swap
    w1 ^= w2;
    w2 ^= w1;
    w1 ^= w2;


    // one's complement
    short int w4 = 2;
    printf("%d\n", ~(w4)); // -3
                           //



    return 0;
} 
