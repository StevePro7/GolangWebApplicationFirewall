#include <stdio.h>
#include <stdlib.h>

//struct ImgInfo
//{
//    char *imgPath;
//};

struct CoreRuleSetErrorObject
{
    char *error_message;
};

void InitializeModSecurity();

void LoadModSecurityCoreRuleSet( struct CoreRuleSetErrorObject *coreRuleSetErrorObject, char *file );

//void LoadModSecurityCoreRuleSetX( struct ImgInfo *imgInfo, char *file );