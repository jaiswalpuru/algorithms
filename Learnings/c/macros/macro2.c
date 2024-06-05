#include <stdio.h>
#include <stdlib.h>

#define DeclareSort(prefix, type) \
    static int \
    _DeclareSort_ ## prefix ## _Compare(const void *a, const void *b) \
{\
    const type *aa; \
    const type *bb; \
    aa = a; \
    bb = b; \
    if (aa < bb) return -1; \
    else if (bb < aa) return 1; \
    return 0; \
} \
\
void prefix ## _sort(type *a, int n) \
{ \
    qsort(a, sizeof(type), n, _DeclareSort_ ## prefix ## _Compare); \
}

// below must appear outside of function and has no trailing spaces
DeclareSort(int, int)

int main(int argc, char **argv)
{
    
    // here we can declare array and sort it
    return 0;
}
