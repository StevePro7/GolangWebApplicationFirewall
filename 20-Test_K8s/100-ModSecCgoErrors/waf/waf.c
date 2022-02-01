#include "waf.h"
#include "modsecurity/modsecurity.h"
#include "modsecurity/rules_set.h"
#include "modsecurity/transaction.h"
#include "modsecurity/intervention.h"
#include <stdio.h>

ModSecurity *modsec = NULL;
RulesSet *rules = NULL;

// General public APIs.
void InitializeModSecurity()
{
    modsec = msc_init();
    //msc_set_log_cb( modsec, CModSecurityLoggingCallback );
    rules = msc_create_rules_set();
}

const char* LoadModSecurityCoreRuleSet( char *file )
{
    const char *error = NULL;
    char *error_message = NULL;
//    if ( modsec == NULL )
//    {
//        initializeModSecurityImpl();
//    }

    msc_rules_add_file( rules, file, &error );
    if ( error != NULL )
    {
        error_message = (char *)error;
    }

    error = NULL;
    return error_message;
}