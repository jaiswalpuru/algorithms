#include <stdio.h>
#include <stdlib.h>

int main(int argc, char **argv)
{
    int num = 0;
    int n = 0;

    printf("Enter a number : ");
    scanf("%d", &num);
    printf("\nEnter the nth bit to check and set : ");
    scanf("%d", &n);

    
    printf("\nNumber before setting the bit : %d\n", num);
    if (((num >> n) & 1) == 1) {
        printf("%d bit is already set\n", n);
    } else {
        num |= (1 << n); 
    }
    printf("number after setting the bit : %d\n", num);

    return 0;  
}
