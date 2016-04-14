#include <stdio.h>
#include <string.h>
#include <ctype.h>
#include <stdlib.h>

#include <math.h>
#include <time.h>

#define OK 1
#define ERROR 0
#define TRUE  1
#define FALSE 0

#define MAXSIZE 20

typedef int Status;

// Example element type
typedef int ElemType;

Status visit(ElemType c)
{
  printf("%d   ", c);
  return OK;
}

typedef struct Node
{
  ElemType data;
  struct Node *next;
}Node;

typedef struct Node *LinkList;

// Init sequence list
Status InitList(LinkList *L)
{
  *L = (LinkList)malloc(sizeof(Node));
  if (!(*L))
    return ERROR;
  (*L)->next = NULL;

  return OK;
}

// Is list empty
Status ListEmpty(LinkList L)
{
  if (L->next)
    return FALSE;

  return TRUE;
}

// Clear list
Status ClearList(LinkList *L)
{
  LinkList p, q;
  p = (*L)->next;
  while(p)
  {
    q = p->next;
    free(p);
    p = q;
  }

  (*L)->next = NULL;
  return OK;
}

// Get list length
int ListLength(LinkList L)
{
  int i = 0;
  LinkList p = L->next;
  while(p)
  {
    i++;
    p = p->next;
  }

  return i;
}


// Get the assign element
Status GetElem(LinkList L, int i, ElemType *e)
{
  int j;
  LinkList p;
  p = L->next;
  j = 1;
  while (p &&j < i)
  {
    p = p->next;
    ++j;
  }

  if ( !p || j > 1)
    return ERROR;

  *e = p->data;
  return OK;
}

// Locate element failed return 0
int LocateElem(LinkList L, ElemType e)
{
  int i = 0;
  LinkList p = L->next;
  while(p)
  {
    i++;
    if(p->data == e)
      return i;

    p = p->next;
  }

  return 0;
}

// insert element to list
Status ListInsert(LinkList *L, int i, ElemType e)
{
  int j;
  LinkList p, s;
  p = *L;
  j = 1;

  // find the locate i element
  while(p && j < i)
  {
    p = p->next;
    ++j;
  }

  if (!p || j > i)
    return ERROR;

  s = (LinkList)malloc(sizeof(Node));
  s->data = e;
  s->next = p->next;
  p->next = s;
  return OK;
}

// Delete the i element
Status ListDelete(LinkList *L, int i, ElemType *e)
{
  int j;
  LinkList p, q;
  p = *L;
  j = 1;

  // find the locate i element
  while(p && j < i)
  {
    p = p->next;
    ++j;
  }

  if (!(p->next) || j > i)
  {
    return ERROR;
  }

  q = p->next;
  p->next = q->next;
  *e = q->data;
  free(q);
  return OK;
}

// Print all element in list
Status ListTraverse(LinkList L)
{
  LinkList p = L->next;
  while(p)
  {
    visit(p->data);
    p = p->next;
  }
  printf("\n");
  return OK;
}

// Create list example insert in the head
void CreateListHead(LinkList *L, int n)
{
  LinkList p;
  int i;
  srand(time(0));
  *L = (LinkList)malloc(sizeof(Node));
  (*L)->next = NULL;
  for (i = 0; i < n; i++)
  {
    p = (LinkList)malloc(sizeof(Node));
    p->data = rand()%100+1;
    p->next = (*L)->next;
    (*L)->next = p;
  }
}

// Create list example insert in the head
void CreateListTail(LinkList *L, int n)
{
  LinkList p, r;
  int i;
  srand(time(0));
  *L = (LinkList)malloc(sizeof(Node));
  r = *L;
  for (i = 0; i < n; i++)
  {
    p = (Node *)malloc(sizeof(Node));
    p->data = rand()%100+1;
    r->next = p;
    r = p;
  }

  r->next = NULL;
  return;
}

int main()
{
  LinkList L;
  ElemType e;
  Status i;
  int j, k;

  i = InitList(&L);
  printf("After init ListLength(L)=%d \n", ListLength(L));

  for (j = 1; j <= 5; j++)
  {
    i = ListInsert(&L, 1, j);
  }
  printf("After insert 1~5 in head of L, L.data=");
  ListTraverse(L);

  printf("ListLength(L)=%d \n", ListLength(L));

  i = ListEmpty(L);
  printf("Is L empty i=%d (1:true 0:false) \n", i);

  i = ClearList(&L);
  printf("After clear list ListLength(L)=%d \n", ListLength(L));
  i = ListEmpty(L);
  printf("Is L empty i=%d (1:true 0:false) \n", i);

  for(j = 1; j <= 10; j++)
  {
    ListInsert(&L, j, j);
  }
  printf("After insert 1~10 in the tail of list L.data=");
  ListTraverse(L);

  printf("ListLength(L) = %d \n", ListLength(L));

  ListInsert(&L, 1, 0);
  printf("After insert 0 in the head of list. L.data= \n");
  ListTraverse(L);
  printf("ListLength(L) = %d \n", ListLength(L));

  GetElem(L, 5, &e);
  printf("The 5 element is %d\n", e);

  for(j = 3; j <=12; j++)
  {
    k = LocateElem(L, j);
    if (k)
      printf("The %d element is %d\n", k, j);
    else
      printf("No element value is %d\n", j);
  }

  k = ListLength(L);
  for(j = k+1; j >= k; j--)
  {
    i = ListDelete(&L, j, &e);
    if (i == ERROR)
      printf("Failed to delete the %d element \n", j);
    else
      printf("Delete the %d element value is %d \n", j, e);
  }

  printf("Print all L element\n");
  ListTraverse(L);

  j = 5;
  ListDelete(&L, j, &e);
  printf("Delete the %d element value is %d", j, e);

  printf("Print all L element\n");
  ListTraverse(L);

  i=ClearList(&L);
  printf("After clear list ListLength(L)=%d \n", ListLength(L));

  CreateListHead(&L, 20);
  printf("Use insert element fron head to create list: ");
  ListTraverse(L);

  i = ClearList(&L);
  printf("After clear list ListLength(L)=%d \n", ListLength(L));

  CreateListTail(&L, 20);
  printf("Use insert element fron tail to create list: ");
  ListTraverse(L);

  return 0;
}
