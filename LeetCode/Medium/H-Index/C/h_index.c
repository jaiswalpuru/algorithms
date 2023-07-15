int cmp(const void* a, const void* b) {
    return (*(int*)a-*(int*)b);
}

int hIndex(int* citations, int citationsSize) {
    qsort(citations, citationsSize, sizeof(int), cmp);
    int i=0;
    while(i < citationsSize && citations[citationsSize-1-i] > i) i++;
    return i;
}
