#include <stdio.h>
#include <stdlib.h>


struct ImgInfo
{
    char *imgPath;
};

void InitializeModSecurity();

void LoadModSecurityCoreRuleSet( struct ImgInfo *imgInfo, char *file );
