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

extern void call_num(ModSecNumCb callback)
{
    char *str = "number";
    printf("C  call_text beg '%s'\n", str);
    callback(17);
    printf("C  call_text end '%s'\n", str);
}
extern void pass_GoNum(void)
{
    printf("C  pass_GoNum beg\n");
    call_num(&GoNum);
    printf("C  pass_GoNum end\n");
}


void call_text(ModSecLogCb callback)
{
    char *str = "bob";
    printf("C  call_text beg '%s'\n", str);
    //callback("foo", "bar");
    callback(str);
    //callback(7);
    printf("C  call_text end '%s'\n", str);
}
void pass_GoText(void)
{
    printf("C  pass_GoText beg\n");
    call_text(&GoText);
    printf("C  pass_GoText end\n");
}

void call_const(ModSecConstCb callback)
{
    char *str = "hello";
    const char *buf = "world";
    printf("C  call_const beg '%s'\n", str);
    callback(str, buf);
    printf("C  call_const end '%s'\n", str);
}
 void pass_GoConst(void)
 {
    printf("C  pass_GoConst beg\n");
    call_const(&GoConst);
    printf("C  pass_GoConst end\n");
 }