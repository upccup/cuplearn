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
typedef int QElemType;

// 循环队列的顺序储存结构
typedef struct {
  QElemType data[MAXSIZE];
  int front; // 头指针
  int rear; // 尾指针 若队列不为空, 指向列尾元素的下一个位置
}SqQueue;


Status visit(QElemType c)
{
  printf("%d ", c);
  return OK;
}

//初始化一个空队列
Status InitQueue(SqQueue *Q)
{
  Q->front = 0;
  Q->rear = 0;
  return OK;
}

// 清空队列
Status ClearQueue(SqQueue *Q)
{
  Q->front=Q->rear=0;
  return OK;
}

// 判读队列是否为空
Status QueueEmpty(SqQueue Q)
{
  if (Q.front == Q.rear)
    return TRUE;

  return FALSE;
}

//获取队列中元素的个数, 即队列长度
int QueueLength(SqQueue Q)
{
  return (Q.rear - Q.front)%MAXSIZE;
}


// 获取队列的头元素, 并返回状态
Status GetHead(SqQueue Q, QElemType *e)
{
  if (Q.rear == Q.front)
    return ERROR;

  *e = Q.data[Q.front];
  return OK;
}

// 在队列尾部插入新的元素作为队尾
Status EnQueue(SqQueue *Q, QElemType e)
{
  // 队列是否已满
  if ((Q->rear+1)%MAXSIZE == Q->front)
    return ERROR;


  Q->data[Q->rear] = e;
  Q->rear = (Q->rear + 1)%MAXSIZE;

  return OK;
}

// 删除队列的首元素, 并获取被删除的值
Status DeQueue(SqQueue *Q, QElemType *e)
{
  if (Q->rear == Q->front)
    return ERROR;

  *e = Q->data[Q->front];
  Q->front = (Q->front + 1) % MAXSIZE;
  return OK;
}


// 遍历队列
Status QueueTraverse(SqQueue Q)
{
  int i;
  i = Q.front;

  while((i+Q.front) != Q.rear)
  {
    visit(Q.data[i]);
    i = (i +1)%MAXSIZE;
  }

  printf("\n");
  return OK;
}

int main()
{
  Status j;
  int i=0, l;
  QElemType d;
  SqQueue Q;

  InitQueue(&Q);
  printf("初始化队列后 队列是否为空%d  1:空 0:非空 \n", QueueEmpty(Q));

  printf("请输入整型队列元素(不超过%d个), -1为提前结束标识符 \n", MAXSIZE-1);

  do
  {
    scanf("%d", &d);
    if (d == -1)
      break;
    EnQueue(&Q, d);
    i++;
  }while(i < MAXSIZE-1);

  printf("队列长度为: %d \n", QueueLength(Q));

  printf("队列是否为空%d  1:空 0:非空 \n", QueueEmpty(Q));
  printf("连续%d次从队首删除元素, 从队尾插入元素: \n", MAXSIZE);

  for(l = 1; l <= MAXSIZE; l++)
  {
    DeQueue(&Q, &d);
    printf("删除的元素是: %d 插入的元素是 %d \n", d, l+100);

    d = l + 100;
    EnQueue(&Q, d);
  }

  printf("现在队列中的元素为: \n");
  QueueTraverse(Q);

  printf("共向队尾插入了 %d 个元素 \n", i+MAXSIZE);
  l = QueueLength(Q);
  if (l-2 > 0)
    printf("现在由队头开始删除 %d 个元素\n", l-2);

  while(QueueLength(Q) > 2)
  {
    DeQueue(&Q, &d);
    printf("删除的元素为 %d \n", d);
  }

  j = GetHead(Q, &d);
  if (j)
    printf("现在队首的元素为 %d \n", d);

  ClearQueue(&Q);
  printf("清空队列后队列是否为空%d  1:空 0:非空 \n", QueueEmpty(Q));
  return 0;
}


