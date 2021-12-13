#include <stdio.h>
#include "_cgo_export.h"
#include "callback.h"

void call_add(adder func) {
    printf("C  Callback beg\n");
    int i = func(4, 3);
    printf("C  Callback end [%d]\n", i);
}

void pass_GoAdd(void)
{
    printf("C  pass_GoAdd beg\n");
    call_add(&GoAdd);
    printf("C  pass_GoAdd end\n");
}