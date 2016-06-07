#include "../include/apue.h"

int main()
{
    char *ptr;
    size_t size;

    ptr = path_alloc(&size);
    if  (getcwd(ptr, size)  == NULL) {
        err_sys("get pwd failed");
    }

    printf("cwd  = %s \n", ptr);

    if (chdir("/usr/lib/zsh/5.0.8/zsh/") < 0) {
        err_sys("chdir failed");
    }

    if  (getcwd(ptr, size)  == NULL) {
        err_sys("get pwd failed");
    }

    printf("cwd  = %s \n", ptr);
    exit(0);
}
