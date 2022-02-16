#include "waf.h"
//#include "_cgo_export.h"
#include "modsecurity/modsecurity.h"
#include "modsecurity/rules_set.h"
#include "modsecurity/transaction.h"
#include "modsecurity/intervention.h"

ModSecurity *modsec = NULL;
RulesSet *rules = NULL;

static void initializeModSecurityImpl()
{
    modsec = msc_init();
//    msc_set_log_cb(modsec, MyFuncPtr);  // TODO
    rules = msc_create_rules_set();
}

int AddRequestHeaders(char **reqHeaderKeys, char **reqHeaderVals, int reqHeaderSize, char *reqBodyText, int reqBodySize)
{
     int retVal = 0;

    Transaction *transaction = NULL;

    const char *reqHeaderKey;
    const char *reqHeaderVal;

    int index = 0;
    if (modsec == NULL)
    {
        initializeModSecurityImpl();
    }


    transaction = msc_new_transaction(modsec, rules, NULL);

    // Request Headers.
    fprintf(stdout, "Beg\n");
    for( index = 0; index < reqHeaderSize; index++ )
    {
        reqHeaderKey = reqHeaderKeys[ index ];
        reqHeaderVal = reqHeaderVals[ index ];

        fprintf(stdout, "Key : '%s'\n", reqHeaderKey);
        fprintf(stdout, "Val : '%s'\n", reqHeaderVal);

        retVal = msc_add_request_header(transaction, reqHeaderKey, reqHeaderVal);
        fprintf(stdout, "SGB : [%d]\n", retVal);

        if ( !retVal )
        {
            //retVal = -5;  //TODO
            goto out;
        }
    }
    fprintf(stdout, "end\n");


    // Request Body.
    fprintf(stdout, "Body : '%s'\n", reqBodyText);
    fprintf(stdout, "size : [%d]\n", reqBodySize);

out:
    msc_transaction_cleanup(transaction);
    return 6;
}

// Helper functions to store all core rule set file names in memory.
char **makeCharArray(int size)
{
    return calloc(sizeof(char*), size);
}
void freeCharArray(char **array, int size)
{
    int index;
    for (index = 0; index < size; index++)
    {
        free(array[index]);
    }

    free(array);
}
void setArrayString(char **array, char *input, int index)
{
    array[index] = input;
}
