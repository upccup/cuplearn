#include "../include/apue.h"


int main()
{
  printf("Hello world from process Id %ld \n", (long)getpid());
  exit(0);
}
