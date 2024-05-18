#include <stdio.h>
#include <stdlib.h>
#include <string.h>

//#define BUF 256

int has_ch(char ch, const char *line);

int main(int argc, char **argv)
{
    FILE *fp = NULL;
    char ch = '\0';
    size_t BUF = 256;
    char line[BUF];
    char *b = line;

    if (argc != 3) {
        printf("invalid cmd args");
        exit(EXIT_FAILURE);
    } 

    ch = argv[1][0];
    
    if ((fp = fopen(argv[2], "r")) == NULL) {
        printf("cant read  file");
        exit(EXIT_FAILURE);
    }

    ssize_t chars = 0;
    while((chars = getline(&b, &BUF, fp)) > 0) {
        if (has_ch(ch, line)) {
            fputs(line, stdout);   
        }
       // printf("%zu\n", chars);
    }

    fclose(fp);

    return 0;
}

int has_ch(char ch, const char *line)
{
    while(*line) {
        if (ch == *line++) {
            return 1;
        }
    }
    return 0;
}
