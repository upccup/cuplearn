#include "../include/apue.h"
#include <fcntl.h>

char buf1[] = "abcdefghij";
char buf2[] = "ABCDEFGHIJ";

int main()
{
  int fd;

  if((fd = creat("file.hole", FILE_MODE)) < 0) {
    err_sys("Create error");
  }

  if(write(fd, buf1, 10) != 10) {
    err_sys("buf1 write error");
  }
  // offset now = 10

  if(lseek(fd, 16380, SEEK_SET) == -1){
    err_sys("lseek error");
  }
  // offset now = 16384

  if (write(fd, buf2, 10) != 10) {
    err_sys("buf2 write error");
  }
  // offset now is 16394

  int noHoleFd;
  if((noHoleFd = creat("file.nohole", FILE_MODE)) < 0) {
    err_sys("Create error");
  }

  int charNum = 0;
  while(charNum < 16390) {
    if (write(noHoleFd, buf2, 10) != 10) {
      err_sys("buf2 write error");
    }
    charNum +=10;
  }

  exit(0);
}
