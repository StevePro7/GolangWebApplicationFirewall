#include "waf.h"
#include "_cgo_export.h"
#include "modsecurity/modsecurity.h"
#include "modsecurity/rules_set.h"
#include "modsecurity/transaction.h"
#include "modsecurity/intervention.h"
#include <stdio.h>

ModSecurity *modsec = NULL;
RulesSet *rules = NULL;

// Private helper function to initialize ModSecurity.
static void initializeModSecurityImpl();

// ModSecurity logging callback infrastructure APIs.
void InvokeModSecurityLoggingCallback( ModSecurityLoggingCallbackFunctionPointer func, char *payload )
{
    // Invoke Golang callback with ModSecurity logging payload i.e. C => Go code invocation.
    func( payload );
}
// Function prototype must match modsecurity.cc ModSecLogCb callback signature.
void CModSecurityLoggingCallback( void *referenceAPI, const void *ruleMessage )
{
    // Remove constness and coerce to char* to be compatible with Golang API.
    char *payload = (char *)ruleMessage;
    InvokeModSecurityLoggingCallback( &GoModSecurityLoggingCallback, payload );
}

// General public APIs.
void InitializeModSecurity()
{
    ModSecurity *modsec = NULL;
    RulesSet *rules = NULL;

    initializeModSecurityImpl();

}
static void initializeModSecurityImpl()
{
    modsec = msc_init();
    msc_set_log_cb( modsec, CModSecurityLoggingCallback );
    rules = msc_create_rules_set();
}

const char* LoadModSecurityCoreRuleSet2( char *file )
{
    const char *error = NULL;
    char *error_message = NULL;
    if ( modsec == NULL )
    {
        initializeModSecurityImpl();
    }

    msc_rules_add_file( rules, file, &error );
    if ( error != NULL )
    {
        error_message = (char *)error;
    }

    error = NULL;
    return error_message;
}


void CleanupModSecurity()
{
    if ( rules != NULL )
    {
        msc_rules_cleanup( rules );
    }
    if ( modsec != NULL )
    {
        msc_cleanup( modsec );
    }

    rules = NULL;
    modsec = NULL;
}
