#include <stdio.h>

int Fib(int i)
{
  if (i < 2)
    return i ==0 ? 0 : 1;

  return Fib(i-1) + Fib(i-2);
}

int main()
{
  int i;
  int a[40];
  printf("迭代显示斐波那契数列: \n");
  a[0] = 0;
  a[1] = 1;

  for(i =2; i < 40; i++)
  {
    a[i] = a[i-1] +a[i-2];
    printf("%d ", a[i]);
  }

  printf("\n");

  printf("递归显示斐波那契数列: \n");
  for(i = 0; i < 40; i++)
  {
    printf("%d ", Fib(i));
  }

  return 0;
}
