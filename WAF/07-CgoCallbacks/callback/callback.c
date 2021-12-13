#include <stdio.h>
#include "calladd.h"

void call_add(adder func) {
    printf("C  Callback beg\n");
    int i = func(2, 4);
    printf("C  Callback end [%d]\n", i);
}