#include <stdio.h>

typedef int(*adder)(int, int);

void call_add(adder)
{
    printf("Calling adder\n");
    int i = func(3, 4);
    printf("Adder returned %d\n", i);
}


void pass_GoAdd(void) {
    call_add(&GoAdd);
}