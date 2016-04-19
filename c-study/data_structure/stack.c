#include <stdio.h>
#include <stdlib.h>

#include <math.h>
#include <time.h>


#define OK 1
#define ERROR 0
#define TRUE 1
#define FALSE 0
#define MAXSIZE 20 /*储存空间的初始分配量*/

typedef int Status;
typedef int SElemType;


// 顺序栈结构
typedef struct
{
  SElemType data[MAXSIZE];
  int top; //用于栈顶指针
}SqStack;



Status visit(SElemType c)
{
  printf("%d  ", c);
  return OK;
}


//构建一个空栈S
Status InitStack(SqStack *S)
{
  S->top = -1;
  return OK;
}

// 把S置为空栈
Status ClearStack(SqStack *S)
{
  S->top = -1;
  return OK;
}

// 判断S是否为空栈
Status StackEmpty(SqStack S)
{
  if (S.top == -1)
    return TRUE;

  return FALSE;
}

//获取栈S的长度, 即栈的长度
int StackLength(SqStack S)
{
  return S.top+1;
}


//若栈不为空返回栈顶元素, 并返回OK, 否则返回ERROR
Status GetTop(SqStack S, SElemType *e)
{
  if (S.top == -1)
    return ERROR;

  *e = S.data[S.top];
  return OK;
}

// 插入元素e 为新的栈顶元素
Status Push(SqStack *S, SElemType e)
{
  // 判断栈是否已满
  if (S->top == MAXSIZE - 1)
    return ERROR;

  S->top++;
  S->data[S->top] = e;
  return OK;
}


// 若栈不为空, 删除掉S的栈顶元素,并返回删除的值和OK, 否则返回ERROR
Status Pop(SqStack *S, SElemType *e)
{
  // 判断是否为空栈
  if (S->top == -1)
    return ERROR;

  *e = S->data[S->top];
  S->top--;
  return OK;
}


// 从栈底到栈顶一次显示所有元素
Status StackTraverse(SqStack S)
{
  int i;
  i = 0;

  while(i <= S.top)
  {
    visit(S.data[i++]);
  }

  printf("\n");
  return OK;
}

int main()
{
  int j;
  SqStack s;
  int e;

  if (InitStack(&s) == OK)
  {
    for(j = 1; j<=10; j++)
    {
      Push(&s, j);
    }
  }

  printf("栈中元素依次为:  ");
  StackTraverse(s);

  Pop(&s, &e);
  printf("弹出的栈顶元素为 e=%d \n", e);

  printf("栈是否为空: %d (1:空 0:非空) \n", StackEmpty(s));

  GetTop(s, &e);
  printf("栈顶元素 e=%d 栈的长度为 %d \n", e, StackLength(s));

  ClearStack(&s);
  printf("清空栈之后 栈是否为空: %d (1:空 0:非空) \n", StackEmpty(s));

  return 0;
}
