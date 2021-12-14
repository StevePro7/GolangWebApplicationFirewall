
int LoadModSecurityCoreRuleSet(char **array, int size);
int ProcessHttpRequest(char *uri, char *http_method, char *http_protocol, char *http_version, char *client_link, int client_port, char *server_link, int server_port);

////void MyFuncPtr(void *foo, const void *bar)
////{
////    fprintf(stderr, "MyFuncPtr! start\n");
////    //fprintf(stderr, "MyFuncPtr '%s'\n", foo);
////    //fprintf(stderr, "MyFuncPtr '%s' [SPLAT!]\n", bar);
////    fprintf(stderr, "MyFuncPtr -end-\n");
////}
//
void InitializeModSecurity();

// Helper functions to store all core rule set file names in memory.
char **makeCharArray(int size);
void freeCharArray(char **array, int size);
void setArrayString(char **array, char *filename, int index);
