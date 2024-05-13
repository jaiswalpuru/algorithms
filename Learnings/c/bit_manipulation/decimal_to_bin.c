#include <stdio.h>
#include <stdlib.h>

typedef long long ll;

ll dec_to_bin(int dec);

int main(int argc, char **argv)
{
    int dec;

    printf("Enter a number : ");
    scanf("%d", &dec);
    printf("binary number : %lld", dec_to_bin(dec));

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
