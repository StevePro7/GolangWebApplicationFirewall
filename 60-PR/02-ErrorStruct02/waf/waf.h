#include <stdio.h>
#include <stdlib.h>


struct ImgInfo
{
    char *imgPath;
};

void InitializeModSecurity();

void LoadModSecurityCoreRuleSet(char *file_name);

void LoadModSecurityCoreRuleSet2(struct ImgInfo *imgInfo, char *file );

void printStruct(struct ImgInfo *imgInfo);
