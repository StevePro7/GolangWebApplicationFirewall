#include <stdio.h>
#include <stdlib.h>


struct ImgInfo
{
    char *imgPath;
};

void InitializeModSecurity();

void LoadModSecurityCoreRuleSet(char *file_name);

void printStruct(struct ImgInfo *imgInfo);
