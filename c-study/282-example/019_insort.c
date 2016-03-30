#include <stdio.h>

/* 直接插入排序 */

void insort(int s[], int n)
{
  int i, j;

  for(i =2; i <= n; i++)
  {
    s[0] = s[i];
    j = i -1;
    while(s[0] < s[j])
    {
      s[j+1] = s[j];
      j--;
    }

    s[j+1] = s[0];
  }
}

int main()
{
  int a[11], i;
  printf("Please enter 10 integer number \n");
  for(i =1; i <11; i++)
  {
    scanf("%d", &a[i]);
  }

  printf("Original sequence: \n");
  for(i = 1; i < 11; i ++)
  {
    printf("%5d", a[i]);
  }
  printf("\n");

  insort(a, 10);

  printf("After sort sequence: \n");

  for(i=1; i<11; i++)
  {
    printf("%5d", a[i]);
  }

  printf("\n");
}

