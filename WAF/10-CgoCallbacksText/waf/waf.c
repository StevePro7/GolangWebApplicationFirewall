#include <stdio.h>
#include "_cgo_export.h"
#include "waf.h"

void call_add(adder func)
{
    printf("C  Callback beg\n");
    int i = func(7, 3);
    printf("C  Callback end [%d]\n", i);
}
void pass_GoAdd(void)
{
    printf("C  pass_GoAdd beg\n");
    call_add(&GoAdd);
    printf("C  pass_GoAdd end\n");
}

void call_text(ModSecLogCb callback)
{
}
void pass_GoText(void)
{
    printf("C  pass_GoText beg\n");
    call_text(&GoText);
    printf("C  pass_GoText end\n");
}
