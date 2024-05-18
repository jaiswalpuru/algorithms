#include <stdio.h>
#include <ctype.h>
#include <stdlib.h>

int main(int argc, char **argv) 
{
    FILE *fp = NULL;
    int words = 0;
    int chars = 0;
    char ch = '\0';
    char s[100];

    if (argc > 2) {
        printf("Too many params");
        exit (1);
    }

    if (argc == 2) {
        fp = fopen(argv[1], "rw");
        if (fp == NULL) {
            printf("error while reading file\n");
            exit (1);
        }

        while (ch != EOF) {
            ch = getc(fp);
            if (isspace(ch)) ++words;
            else chars++;
        }
        fclose(fp);
    } else {
        printf("Enter a string : ");
        while((ch = getchar()) != EOF) {
            if (isspace(ch)) ++words;
            else chars++;
        }

    }
    printf("words : %d\n", words+1);
    printf("chars : %d\n", chars);


    return 0;
}

