#include "waf.h"
#include "_cgo_export.h"
#include "modsecurity/modsecurity.h"
#include "modsecurity/rules_set.h"

ModSecurity *modsec = NULL;
RulesSet *rules = NULL;

// Private helper function to initialize ModSecurity.
static void initializeModSecurityImpl();

// Function prototype must match modsecurity.cc ModSecLogCb callback signature.
void CModSecurityLoggingCallback( void *referenceAPI, const void *ruleMessage )
{
    // Remove constness and coerce to char* to be compatible with Golang API.
    char *payload = (char *)ruleMessage;
    GoModSecurityLoggingCallback( payload );
}

// General public APIs.
void InitializeModSecurity()
{
    initializeModSecurityImpl();
}
static void initializeModSecurityImpl()
{
    modsec = msc_init();
//    msc_set_log_cb( modsec, null );
    rules = msc_create_rules_set();
}

void LoadModSecurityCoreRuleSet( struct CoreRuleSetErrorObject *coreRuleSetErrorObject, char *file )
{
    const char *error = NULL;
    if ( modsec == NULL )
    {
        initializeModSecurityImpl();
    }

    msc_rules_add_file( rules, file, &error );
    if ( error != NULL )
    {
        coreRuleSetErrorObject->error_message = (char *)error;
    }

    error = NULL;
}