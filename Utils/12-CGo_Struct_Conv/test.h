#include <stdio.h>
#include <stdlib.h>
struct x
{
    int y;
    int z;
};

struct bob
{
    int idx;
};

struct ImgInfo
{
    char *imgPath;
};
void test(struct bob *b)
{
    fprintf(stderr, "bob1 idx = %d\n", b->idx);
    b->idx = 14;
    fprintf(stderr, "bob2 idx = %d\n", b->idx);
}

void printStruct(struct ImgInfo *imgInfo)
{
    fprintf(stdout, "imgPath1 = %s\n", imgInfo->imgPath);
    imgInfo->imgPath = "change this value";
    fprintf(stdout, "imgPath2 = %s\n", imgInfo->imgPath);
}
//void fishing(struct fush *xx)
//{
//    fprintf(stderr, "fush1 STR = %s\n", xx->str);
//    //b->idx = 14;
//    fprintf(stderr, "fush2 STR = %s\n", xx->str);
//}

int sum(struct x a)
{
    return a.y + a.z;
}

int MyCAdd(int x, int y)
{
    return x + y;
}
int MyCInit()
{
    return 72;
}