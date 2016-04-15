#include <stdio.h>
#include <ctype.h>
#include <string.h>
#include <stdlib.h>
#include <math.h>
#include <time.h>


#define OK 1
#define ERROR 0
#define TRUE 1
#define FALSE 0

// Capacity inin space
#define MAXSIZE 1000

typedef int Status;
typedef char ElemType;

Status visit(ElemType c)
{
  printf("%c   ", c);
  return OK;
}

typedef struct
{
  ElemType data;
  // cursor, if cursor is 0 point to null
  int cur;
} Component, StaticLinkList[MAXSIZE];

// 将一维数组sapce 中各分量链成一个备用链表吗space[0].cur 为头指针
// "0" 表示空指针
Status InitList(StaticLinkList space)
{
  int i;
  for (i = 0; i < MAXSIZE-1 ; i++) {
    space[i].cur = i+1;
  }

  //目前静态链表为空, 最后一个元素的cur为0
  space[MAXSIZE-1].cur = 0;
  return OK;
}


// 若备用空间链表非空, 则返回分配的节点下标, 否则返回0
int Malloc_SSL(StaticLinkList space)
{
  // 当前数组第一个元素的cur存的值就是要返回的第一个备用空闲的下标
  int i = space[0].cur;

  // 由于要拿出一个分量来使用了, 所以就得把它的下一个分量用来做备用
  if (i)
    space[0].cur = space[i].cur;

  return i;
}


// 将下标为k 的空闲节点回收到备用链表
void Free_SSL(StaticLinkList space, int k)
{
  // 把第一个元素的cur值赋给要删除的分量的cur
  // 把要删除的分量下标赋值给第一个元素的cur
  space[k].cur = space[0].cur;
  space[0].cur = k;
}


// 获取L中数据元素的个数
int ListLength(StaticLinkList L)
{
  int j = 0;
  int i = L[MAXSIZE-1].cur;
  while (i)
  {
    i = L[i].cur;
    j++;
  }

  return j;
}

Status ListInsert(StaticLinkList L, int i, ElemType e)
{
  int j, k, l;
  k = MAXSIZE - 1;
  if (i < 1 || i > ListLength(L) + 1)
    return ERROR;

  j = Malloc_SSL(L);
  if (j)
  {
    L[j].data = e;
    for (l = 1; l <= i - 1; l++)
    {
      k = L[k].cur;
    }

    L[j].cur = L[k].cur;
    L[k].cur = j;
    return OK;
  }

  return ERROR;
}


Status ListDelete(StaticLinkList L, int i)
{
  int j, k;
  if (i < 1 || i > ListLength(L))
    return ERROR;

  k = MAXSIZE - 1;
  for (j = 1; j <= i - 1; j++)
  {
    k = L[k].cur;
  }

  j = L[k].cur;
  L[k].cur = L[j].cur;
  Free_SSL(L, j);
  return OK;
}


Status ListTraverse(StaticLinkList L) {
  int j = 0;
  int i = L[MAXSIZE-1].cur;
  while(i)
  {
    visit(L[i].data);
    i = L[i].cur;
    j++;
  }
  printf("\n");
  return OK;
}

int main()
{
  StaticLinkList L;
  Status i;
  i = InitList(L);
  printf("After init L L.length=%d \n", ListLength(L));

  i = ListInsert(L, 1, 'F');
  i = ListInsert(L, 1, 'E');
  i = ListInsert(L, 1, 'D');
  i = ListInsert(L, 1, 'B');
  i = ListInsert(L, 1, 'A');

  printf("After insert FEDBA in the head of L \n L.data= ");
  ListTraverse(L);

  i = ListInsert(L, 3, 'C');
  printf("After insert C between B and D \n L.data= ");
  ListTraverse(L);

  i = ListDelete(L, 1);
  printf("After delete A \n L.data= ");
  ListTraverse(L);

  printf("\n");
  return 0;
}
