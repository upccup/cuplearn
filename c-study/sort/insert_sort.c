#include <stdio.h>

/*
 * 1. 从第一个元素开始改元素可以认为已经被排序
 * 2. 取出下一个元素, 在已经排序的元素序列中从后向前扫描
 * 3. 如果该元素大于新元素(已排序), 将元素移到下一位置
 * 4. 重复步骤3, 直到找到已排序的元素小于或等于新元素的位置
 * 5. 将新元素查到该位置中
 * 6. 重复 2~5
*/

void insert_sort(int a[], int n)
{
  int temp, j;
  for(int i = 1; i < n ; i++)
  {
    temp = a[i];
    for(j = i; j>0 && temp < a[j-1] ; j--)
    {
       a[j] = a[j-1];
    }
    a[j] = temp;
  }
}

// error sort 
void insert_sort2(int a[], int n)
{
  int temp, j;
  for(int i = 1; i < n ; i++)
  {
    temp = a[i];
    for(j = i; j>0 && temp < a[j-1] ; j--)
    {
       continue;
    }
    printf("num: %d, seq: %d \n", temp, j);
    a[i] = a[j];
    a[j] = temp;
  }
}

int main()
{
  int a[10] = {23, 1, 34, 21, 11, 4, 5, 2, 1, 12};
  for(int i = 0; i < 10; i++)
  {
    printf("a[%d]: %d \n", i, a[i]);
  }

  insert_sort(a, 10);
  for(int i = 0; i < 10; i++)
  {
    printf("a[%d]: %d \n", i, a[i]);
  }

  int b[]= {11 ,3 ,1 ,5 ,2 ,23 ,8 ,4 ,1 ,123};
  for(int i = 0; i < 10; i++)
  {
    printf("b[%d]: %d \n", i, b[i]);
  }

  insert_sort2(b, 10);
  for(int i = 0; i < 10; i++)
  {
    printf("b[%d]: %d \n", i, b[i]);
  }
  return 0;
}
