#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

bool is_even(long long *num);
bool is_odd(long long *num);


int main(int argc, char **argv)
{
    FILE *fp = NULL;
    long long *num;

    fp = fopen("numbers.txt", "r");
    if (fp == NULL) {
        printf("can't open the file\n");
        exit(EXIT_FAILURE);
    }
    
    while (fscanf(fp, "%lld", num) == 1) {
        if (is_even(num)) {
            printf("Even number : %lld\n", *num);
        } else if (is_odd(num)) {
            printf("Odd number : %lld\n", *num);
        }
    }

    fclose(fp);
    return 0;
}

bool is_even(long long *num) { return *num%2 == 0 ? true : false; }
bool is_odd(long long *num) { return *num%2 != 0 ? true : false; }
// need prime also but got lazy to write the prime func :_)
