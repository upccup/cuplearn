#include <stdio.h>

/*
 求x^y的最后三位数, x, y是键盘输入的整数
*/


void main()
{
  int i, x, y, z = 1;
  printf("please enter two number x and y (x^y): \n");
  scanf("%d%d", &x, &y);
  for(i = 1; i <= y; i++)
    z = z * x % 1000;

  if (z>=100)
  {
    printf("%d^%d的最后三位是: %d\n", x, y, z);
  }
  else
  {
    printf("%d^%d的最后三位是: 0%d\n", x, y, z);
  }
}
