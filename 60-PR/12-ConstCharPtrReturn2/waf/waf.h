// Important: <stdlib.h> header file required for Golang statements e.g. defer C.free(unsafe.Pointer(API))
#include <stdlib.h>

// ModSecurity logging callback infrastructure APIs.
typedef void( *ModSecurityLoggingCallbackFunctionPointer )( char * );
extern void InvokeModSecurityLoggingCallback( ModSecurityLoggingCallbackFunctionPointer, char * );

void InitializeModSecurity();
const char* LoadModSecurityCoreRuleSet2( char *file );
void CleanupModSecurity();
