#include <stdio.h>
#include <stdlib.h>

struct CoreRuleSetErrorObject
{
    char *error_message;
};

void InitializeModSecurity();

void LoadModSecurityCoreRuleSet( struct CoreRuleSetErrorObject *coreRuleSetErrorObject, char *file );
