#include <stdio.h>
#include <stdlib.h>



int main(int argc, char **argv)
{
#if __STDC__
    printf("implemeting ISO confirming standard\n");
    printf("line number %d\n", __LINE__);
    printf("file name %s\n", __FILE__);
    printf("todays date %s\n", __DATE__);
    printf("time %s\n", __TIME__);
#else
    printf("Non confirming standard\n");
#endif

    return 0;
}
