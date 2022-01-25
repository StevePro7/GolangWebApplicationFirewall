#include <stdio.h>

struct x
{
    int y;
    int z;
};

struct bob
{
    int idx;
};

void test(struct bob *b)
{
    fprintf(stderr, "bob1 idx = %d\n", b->idx);
    b->idx = 7;
    fprintf(stderr, "bob2 idx = %d\n", b->idx);
}

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