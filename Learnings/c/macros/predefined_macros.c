#include <stdio.h>
#include <stdlib.h>

void my_function_name(void) {
    printf("my name is %s", __func__);
}

int main(int argc, char **argv)
{
#if __STDC__
    printf("Implementation if ISO-confirming");
#else
    printf("Implementation is non-Iso confirming");
#endif
    my_function_name();
    return 0;
}
