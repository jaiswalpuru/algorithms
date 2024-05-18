#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX 41

int main(int argc, char **argv)
{
    char string[100];
    char ch[100];

    sprintf(string, "%d, hello this is the number %s", 55, "bye");
    puts(string);

    FILE *fp = NULL;
    if ((fp = fopen("text.txt", "r+")) == NULL) {
        printf("can't open the file\n");
        exit(EXIT_FAILURE);
    }

    for (int i = 0; i < 10; i++) {
        fprintf(fp, "count : %d\n", i);
    }
    fclose(fp);

    if ((fp = fopen("text.txt", "r")) == NULL) {
        printf("can't open the file\n");
        exit(EXIT_FAILURE);
    }

    while (!feof(fp)) {
        fgets(ch, 100, fp);
        printf("%s", ch);
    }
    fclose(fp);

    char buff[255];
    fp = fopen("text.txt", "r");
    while (fscanf(fp, "%s", buff) != EOF) {
        printf("%s", buff);
    }
    fclose(fp);
    
    char words[100];
    if ((fp = fopen("text.txt", "a+")) == NULL)// a+ allows to read and write (append)
    {
        fprintf(stderr, "can't open the file\n");
        exit(EXIT_FAILURE);
    }

    puts("Enter the words to add to the file, press # key to terminate");

    while ((fscanf(stdin, "%40s", words) == 1) && (words[0] != '#')) {
        fprintf(fp, "%s\n", words);
    } 


    puts("File contents");
    rewind(fp); // go back to beginnig of the file
                //

    while (fscanf(fp, "%s", words) == 1) {
        puts(words);
    }

    if (fclose(fp) != 0) exit(EXIT_FAILURE);


    char *str = "Heelo worlds 40";
    char name[10], title[10];
    int age;
    int ret = sscanf(str, "%s %s %d", name, title, &age);
    printf("%s\n%s\n%d\n", name, title, age);

    return 0;
}

