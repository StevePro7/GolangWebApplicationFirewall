#ifndef __CALICO_WAF_H__
#define __CALICO_WAF_H__

#include <stdlib.h>

struct CoreRuleSetErrorObject
{
    char *error_message;
};

void InitializeModSecurity();

void LoadModSecurityCoreRuleSet( struct CoreRuleSetErrorObject *coreRuleSetErrorObject, char *file );

#endif//__CALICO_WAF_H__