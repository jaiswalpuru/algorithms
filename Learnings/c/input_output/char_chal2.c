#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>

void convert_case(FILE *fp, const char *path);

int main(int argc, char **argv)
{
    FILE *fp = NULL;
    char path[30];

    printf("enter file name : ");
    scanf("%s", path);

    fp = fopen(path, "r");
    if (fp == NULL) {
        printf("wrong path\n");
        exit(EXIT_FAILURE);
    }
    
    convert_case(fp, path);

    return 0;
}

void convert_case(FILE *fp, const char *path)
{
    FILE *des = NULL;
    char ch = '\0';

    des = fopen("output.txt", "w");
    if (des == NULL) {
        printf("failed to open file \n");
        exit(EXIT_FAILURE);
    }

    while ((ch = fgetc(fp)) != EOF) {
        if (isupper(ch))
            ch = tolower(ch);
        else if (islower(ch)) 
            ch = toupper(ch);
        fputc(ch, des);
    }

    fclose(fp);
    fclose(des);

    remove(path); // remove the file
    rename("output.txt", path); // rename the output file to original file
}
