#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(int argc, char **argv) 
{
    /*
    char buff[255];
    int ch = '\0';
    char *p = NULL;

    if (fgets(buff, sizeof(buff), stdin)) {
        p = strchr(buff, '\n');
        if (p) {
            *p = '\0';
        }
        else {
            while (((ch = getchar()) != '\n') && !feof(stdin) && !ferror(stdin)); // just dumping rest of characters after line feed
        }
    }*/



    char *buffer = NULL;
    size_t bufsize = 32;
    size_t characters;

    buffer = (char *)malloc(bufsize * (sizeof(char)));
    if (buffer == NULL) exit(0);

    characters = getline(&buffer, &bufsize, stdin);

    printf("%zu characters were read\n ", characters);
    printf("%s\n", buffer);
    puts(buffer);

    char string[40];
    strcpy(string, "Hellow world");
    puts(string);

    // fputs() , writes to stream
    return 0;
}
