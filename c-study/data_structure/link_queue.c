#include <stdio.h>
#include <stdlib.h>

#include <math.h>
#include <time.h>

#define OK 1
#define ERROR 0
#define TRUE 1
#define FALSE 0

typedef int Status;

typedef int QElemType;

typedef struct QNode
{
  QElemType data;
  struct QNode *next;
}QNode, *QueuePtr;


typedef struct
{
  QueuePtr front, rear;
}LinkQueue;


Status visit(QElemType c)
{
  printf("%d ", c);
  return OK;
}


// 构造一个空队列
Status InitQueue(LinkQueue *Q)
{
  Q->front = Q->rear = (QueuePtr)malloc(sizeof(QNode));
  if (!Q->front)
    exit(OVERFLOW);

  Q->front->next = NULL;
  return OK;
}

// 销毁队列
Status Destroy(LinkQueue *Q)
{
  while(Q->front)
  {
    Q->rear = Q->front->next;
    free(Q->front);
    Q->front = Q->rear;
  }
  return OK;
}

//将队列清空为空队列
Status ClearQueue(LinkQueue *Q)
{
  QueuePtr p, q;
  Q->rear = Q->front;
  p = Q->front->next;
  Q->front->next = NULL;

  while(p)
  {
    q = p;
    p = p->next;
    free(q);
  }

  return OK;
}

// 判断队列是否为空
Status QueueEmpty(LinkQueue Q)
{
  if (Q.rear == Q.front)
    return TRUE;

  return FALSE;
}

// 获取队列长度 即队列中元素的个数
int QueueLength(LinkQueue Q)
{
  int i = 0;
  QueuePtr p;
  p = Q.front;
  while (Q.rear != p)
  {
    i++;
    p = p->next;
  }

  return i;
}

// 获取队列的头元素
Status GetHead(LinkQueue Q, QElemType *e)
{
  QueuePtr p;
  if (Q.rear == Q.front)
    return ERROR;

  p = Q.front->next;
  *e = p->data;
  return OK;
}

// 向队列底部插入新的元素e
Status EnQueue(LinkQueue *Q, QElemType e)
{
  QueuePtr s = (QueuePtr)malloc(sizeof(QNode));
  if (!s)
    exit(OVERFLOW); //分配内存失败

  s->data = e;
  s->next = NULL;
  Q->rear->next = s;
  Q->rear = s;
  return OK;
}

// 删除列头元素并返回被删除的值
Status DeQueue(LinkQueue *Q, QElemType *e)
{
  QueuePtr p;
  if (Q->front == Q->rear)
    return ERROR;

  p = Q->front->next;
  *e = p->data;

  Q->front->next = p->next;
  if (Q->rear == p)
    Q->rear = Q->front;

  free(p);
  return OK;
}

// 遍历队列
Status QueueTraverse(LinkQueue Q)
{
  QueuePtr p;

  p = Q.front->next;
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
  int i;
  QElemType d;
  LinkQueue q;
  i = InitQueue(&q);
  if (i)
    printf("构造一个空队列成功\n");

  printf("队列是否为空: %d (1:空, 0:非空) \n", QueueEmpty(q));

  printf("队列长度为%d \n", QueueLength(q));

  EnQueue(&q, 5);
  EnQueue(&q, -5);
  EnQueue(&q, 10);
  printf("插入三个元素 (5, -5, 10) 之后队列长度为 %d \n", QueueLength(q));

  printf("队列是否为空: %d (1:空, 0:非空) \n", QueueEmpty(q));

  printf("队列中的元素一次为: ");
  QueueTraverse(q);

  i = GetHead(q, &d);
  if (i == OK)
    printf("队列首元素为%d \n", d);

  DeQueue(&q, &d);
  printf("删除了列首元素 %d \n", d);

  i = GetHead(q, &d);
  if (i == OK)
    printf("新的队列首元素为%d \n", d);

  ClearQueue(&q);
  printf("清空队列后 q.front=%u, q.rear=%u, q.front->next=%u\n", q.front, q.rear, q.front->next);

  Destroy(&q);
  printf("销毁队列后 q.front=%u, q.rear=%u \n", q.front, q.rear);

  return  0;
}
