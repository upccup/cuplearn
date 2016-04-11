#include <stdio.h>

#include <stdlib.h>
#include <math.h>
#include <time.h>

#define OK 1
#define ERROR 0
#define TRUE 1
#define FALSE 0

// defakut init space allocation
#define MAXSIZE 20

typedef int Status;

// example element type
typedef int ElemType;

Status visit(ElemType c)
{
  printf("%d  ", c);
  return OK;
}

typedef struct
{
  // store element arry
  ElemType data[MAXSIZE];
  // list current length
  int length;
}SqlList;


// init order linear table
Status InitList(SqlList *L)
{
  L->length =0;
  return OK;
}

// judge list is empty
Status ListEmpty(SqlList L)
{
  if (L.length == 0 )
    return TRUE;
  else
    return FALSE;
}

// clear list
Status ClearList(SqlList *L)
{
  L -> length = 0;
  return OK;
}

// get list length
int ListLength(SqlList L)
{
  return L.length;
}

// get element
Status GetElem(SqlList L, int i, ElemType *e)
{
  if (L.length == 0 || i < 1 || i > L.length)
    return FALSE;

  *e = L.data[i-1];
  return OK;
}

// find element
int LocateElem(SqlList L, ElemType e)
{
  int i;
  if (L.length == 0)
    return -0;

  for (i = 0; i < L.length ; i++) {
    if (L.data[i] == e)
    {
      break;
    }
  }

  if ( i >= L.length)
  {
    return 0;
  }

  return i + 1;
}

// insert element
Status ListInsert(SqlList *L, int i, ElemType e)
{
  int k;
  if (L->length == MAXSIZE)
  {
    return ERROR;
  }

  if (i < 1 || i > L->length + 1)
  {
    return ERROR;
  }

  if (i <= L->length)
  {
    for (k = L->length-1; k >= i-1; k--)
    {
      L->data[k+1] = L->data[k];
    }
  }

  L->data[i-1] = e;
  L->length++;

  return OK;
}

// delete element
Status ListDelete(SqlList *L, int i, ElemType *e)
{
  int k;
  if (L->length == 0)
  {
    return ERROR;
  }

  if (i<1 || i>L->length)
  {
    return ERROR;
  }

  *e = L->data[i-1];
  if (i<L->length)
  {
    for (k = i; k < L->length ; k++)
    {
      L->data[k-1] = L->data[k];
    }
  }

  L->length--;
  return OK;
}

// print list
Status ListTraverse(SqlList L)
{
  int i;
  for (i = 0; i < L.length; i++)
  {
    visit(L.data[i]);
  }

  printf("\n");
  return OK;
}

void unionl(SqlList *La, SqlList Lb)
{
  int La_len, Lb_len, i;
  ElemType e;
  La_len = ListLength(*La);
  Lb_len = ListLength(Lb);
  for (i = 1; i <= Lb_len; i++)
  {
    GetElem(Lb, i, &e);
    if(!LocateElem(*La, e))
    {
      ListInsert(La, ++La_len, e);
    }
  }
}

int main()
{
  SqlList L, Lb;
  ElemType e;
  Status i;
  int j, k;
  i = InitList(&L);
  printf("After init L L.length=%d \n", L.length);
  for (j = 1; j <= 5; j++)
  {
    i = ListInsert(&L, 1, j);
  }
  printf("After insert 1~5 in head of list L.data=");
  ListTraverse(L);
  printf("L.length=%d \n", L.length);
  i = ListEmpty(L);
  printf("Is L empty i=%d (1 is yes 0 is no) \n", i);

  i = ClearList(&L);
  printf("After clear L L.length=%d \n", L.length);
  i = ListEmpty(L);
  printf("Is L empty i=%d (1 is yes 0 is no) \n", i);

  for(j = 1; j <= 10; j++)
  {
    ListInsert(&L, j, j);
  }
  printf("After insert 1~10 in the end of list: L.data=");
  ListTraverse(L);

  printf("L.length = %d \n", L.length);

  ListInsert(&L, 1, 0);
  printf("After insert 0 in the head of L: L.data=");
  ListTraverse(L);
  printf("L.length = %d \n", L.length);

  GetElem(L, 5, &e);
  printf("The five element is: %d \n", e);
  for(j = 3; j <= 12; j++)
  {
    k = LocateElem(L, j);
    if (k)
      printf("The %d element value is %d \n", k, j);
    else
      printf("Not found %d element\n", j);
  }

  k = ListLength(L);
  for(j = k+1; j>= k; j--)
  {
    i = ListDelete(&L, j, &e);
    if (i == ERROR)
      printf("Delete the %d element failed \n", j);
    else
      printf("Delete the %d element is %d \n", j, e);
  }

  printf("Print all element of L \n");
  ListTraverse(L);

  j = 5;
  ListDelete(&L, j, &e);
  printf("Delete the %d element is %d \n", j, e);

  printf("Print all element of L \n");
  ListTraverse(L);

  i = InitList(&Lb);
  for(j = 6; j  <=15; j++)
  {
    i = ListInsert(&Lb, 1, j);
  }

  unionl(&L, Lb);

  printf("Print all element of L \n");
  ListTraverse(L);

  return 0;
}


