#include <stdio.h>
#include <stdlib.h>

#include <math.h>
#include <time.h>



#define OK 1
#define ERROR 0
#define TRUE 1
#define FALSE 0
#define MAXSIZE 20

typedef int Status;
typedef int SElemType;

// 链栈结构
typedef struct StackNode
{
  SElemType data;
  struct StackNode *next;
}StackNode, *LinkStackPtr;


typedef struct
{
  LinkStackPtr top;
  int count;
}LinkStack;


Status visit(SElemType c)
{
  printf("%d  \n", c);
  return OK;
}


// 构造一个空栈
Status InitSatack(LinkStack *S)
{
  S->top = (LinkStackPtr)malloc(sizeof(StackNode));
  if (!S->top)
    return ERROR;

  S->top = NULL;
  S->count = 0;
  return OK;
}


// 清空栈
Status ClearStack(LinkStack *S)
{
  LinkStackPtr p, q;
  p = S->top;
  while(p)
  {
    q = p;
    p = p->next;
    free(q);
  }

  S->count = 0;
  return OK;
}

// 判断是否为空栈
Status StackEmpty(LinkStack S)
{
  if (S.count == 0 )
    return TRUE;

  return FALSE;
}


// 获取栈中元素的个数
int StackLength(LinkStack S)
{
  return S.count;
}


// 获取栈顶元素
Status GetTop(LinkStack S, SElemType *e)
{
  if (S.top == NULL)
    return ERROR;

  *e = S.top->data;
  return OK;
 }

// 插入元素e 为新的栈顶元素
Status Push(LinkStack *S, SElemType e)
{
  LinkStackPtr s = (LinkStackPtr)malloc(sizeof(StackNode));
  s->data = e;
  s->next = S->top;
  S->top = s;
  S->count++;
  return OK;
}


// 删除栈顶元素 返回被删除的值
Status Pop(LinkStack *S, SElemType *e)
 {
  if (StackEmpty(*S))
    return ERROR;

  LinkStackPtr p;
  p = S->top;
  S->top = S->top->next;
  *e = p->data;
  free(p);
  S->count--;

  return OK;
 }


Status StackTraverse(LinkStack S)
{
  LinkStackPtr p;
  p = S.top;
  while(p)
  {
    visit(p->data);
    p = p->next;
  }

  printf("\n");
  return OK;
}

int main()
{
  int j;
  LinkStack s;
  int e;
  if (InitSatack(&s) == OK)
  {
    for (j = 1; j<=10; j ++)
    {
      Push(&s, j);
    }
  }

  printf("栈中的元素依次为:   \n");
  StackTraverse(s);
  Pop(&s, &e);
  printf("弹出栈顶元素e=%d  \n", e);
  printf("栈是否为空 %d (1:空, 0:非空) \n", StackEmpty(s));

  GetTop(s, &e);
  printf("栈顶元素 e=%d 栈的长度为%d \n", e, StackLength(s));

  ClearStack(&s);
  printf("清空栈之后 栈是否为空 %d (1:空, 0:非空) \n", StackEmpty(s));
  return 0;
}
