#include "waf.h"


int AddRequestHeaders(char **array, int size)
{
    return 0;
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
