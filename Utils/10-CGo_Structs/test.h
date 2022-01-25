#include <stdio.h>
#include <stdlib.h>

char ch = 'M';
unsigned char uch = 253;
short st = 233;
int i = 257;
long lt = 11112222;
float f = 3.14;
double db = 3.15;
void * p;
char *str = "const string";
char str1[64] = "char array";

struct ImgInfo
{
    char *imgPath;
    int format;
    unsigned
    int width;
    unsigned int height;
};

void printI(void *i)
{
    printf("print i = %d\n", (*(int *)i));
}

void printStruct(struct ImgInfo *imgInfo)
{
    if(!imgInfo)
    {
        fprintf(stderr, "imgInfo is null\n");
        return;
    }

    fprintf(stdout, "imgPath = %s\n", imgInfo->imgPath);
    fprintf(stdout, "format = %d\n", imgInfo->format);
    fprintf(stdout, "width = %d\n", imgInfo->width);
}