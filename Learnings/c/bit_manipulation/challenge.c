#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

typedef long long ll;

ll dec_to_bin(int dec);

int main(int argc, char **argv)
{
    int a = 0, b = 0;
    
    printf("Enter two number\n");
    scanf("%d %d", &a, &b);
    
    printf("one's complement\n");
    int a_oc = ~(a);
    int b_oc = ~(b);
    bool is_a_neg = a_oc > 0 ? false : true;
    bool is_b_neg = b_oc > 0 ? false : true;

    ll a_bin =  dec_to_bin(is_a_neg ? (-1 * a_oc) : a_oc);
    ll b_bin =  dec_to_bin(is_b_neg ? (-1 * b_oc) : b_oc);

    printf("~a = %lld ~b = %lld\n", is_a_neg ? (-1 * a_bin) : a_bin, is_b_neg ? (-1 * b_bin) : b_bin);
    printf("a & b : %lld\n", dec_to_bin(a & b));
    printf("a | b : %lld\n", dec_to_bin(a | b));
    printf("a ^ b : %lld\n", dec_to_bin(a ^ b));
    printf("a << 2 : %lld\n", dec_to_bin(a << 2));


    return 0;
}

ll dec_to_bin(int dec)
{
    ll res = 0;
    int rem = 0;
    int i = 1;

    while (dec)
    {
        rem = dec%2;
        dec /= 2;
        res += rem*i;;
        i *= 10;
    }

    return res;
}
