#include <stdio.h>
#include "_cgo_export.h"
#include "waf.h"

void call_add(adder func)
{
    printf("C  call_add beg\n");
    int i = func(1, 2);
    printf("C  call_add end [%d]\n", i);
}
void pass_GoAdd(void)
{
    printf("C  pass_GoAdd beg\n");
    call_add(&GoAdd);
    printf("C  pass_GoAdd end\n");
}

void call_text(ModSecLogCb callback)
{
    printf("C  call_text beg\n");
    callback("foo", "bar");
    printf("C  call_text end\n");
}
void pass_GoText(void)
{
    printf("C  pass_GoText beg\n");
    call_text(&GoText);
    printf("C  pass_GoText end\n");
}
