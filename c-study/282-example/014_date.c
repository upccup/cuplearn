#include <stdio.h>


int leap(int a)
{
  if((a%4 == 0 && a%100 != 0) || a%400 == 0)
    return 1;
  else
    return 0;
}

int number(int year, int m, int d)
{
  int sum = 0, i, j, k, a[12]=
  {
    31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31
  };
  int b[12] =
  {
    31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31
  };

  if (leap(year) == 1)
    for(i=0;i<m-1;i++)
      sum += b[i];
  else
    for(i=0;i<m-1;i++)
      sum += a[i];

  sum +=d;
  return sum;
}

int main()
{
  int year, month, day, n;
  printf("please enter year month day \n");
  scanf("%d%d%d", &year, &month, &day);

  n = number(year, month, day);
  printf("The%d day\n", n);

  return 0;
}
