#include <stdio.h>
#include <stdlib.h>

#include <math.h>
#include <time.h>

#define OK 1
#define ERROR 0
#define TRUE 1
#define FALSE 0

#define MAXSIZE 100
#define MAX_TREE_SIZE 100 //二叉树最大节点数

typedef int Status;
typedef int TElemType;
typedef TElemType SqBiTree[MAX_TREE_SIZE];

typedef struct
{
  int level, order;  //节点的层和本层序号(按满二叉树计算)
}Position;

TElemType Nil = 0;

Status visit(TElemType c)
{
  printf("%d  ", c);
  return OK;
}


// 构造空二叉树, 因为T是固定数组不会改变, 故不需要&
Status InitBiTree(SqBiTree T)
{
  int i;
  for(i=0; i<MAX_TREE_SIZE; i++)
    T[i] = Nil; //初始值为空

  return OK;
}

// 按层序次序输入二叉树中节点的值(字符型或整型), 构造顺序存储的二叉树
Status CreateBiTree(SqBiTree T)
{
  int i = 0;
  printf("请按层序输入节点的值(整型), 0表示空节点,输入999结束, 节点数<= %d: \n", MAX_TREE_SIZE);
  while(i<10)
  {
    T[i] = i+1;

    if(i != 0 && T[(i+1)/2-1] == Nil && T[i] != Nil) // 此节点不为空无双亲且不是根
    {
      printf("出现无双亲非根节点的节点: %d \n", T[i]);
      exit(ERROR);
    }
    i++;
  }

  while(i < MAX_TREE_SIZE)
  {
    T[i] = Nil; // 将空赋值给T的后面的节点o
    i++;
  }

  return OK;
}

#define ClearBiTree InitBiTree   //在顺序储存结构中 两函数完全一样

// 判断二叉树是否为空
Status BiTreeEmpty(SqBiTree T)
{
  if(T[0] == Nil)
  {
    return TRUE;
  }
  else
  {
    return FALSE;
  }
}


// 获取二叉树的深度
int BiTreeDepth(SqBiTree T)
{
  int i, j = -1;
  for(i=MAX_TREE_SIZE-1; i>=0; i--)
  {
    if(T[i] == Nil)
    {
      break;
    }
  }

  i++;
  do
  {
    j++;
  }while(i>=pow(2,i)); // 计算2的j次幂

  return j;
}

// 获取二叉树T的根
Status Root(SqBiTree T, TElemType *e)
{
  if(BiTreeEmpty(T))
  {
    return ERROR;
  }
  else
  {
    *e = T[0];
    return OK;
  }
}

// 获取处于位置e(层, 本层序号)的节点的值
TElemType Value(SqBiTree T, Position e)
{
  return T[(int)powl(2, e.level-1)+e.order-2];
}

// 给位于位置e的节点赋值
Status Assign(SqBiTree T, Position e, TElemType value)
{
  int i = (int)powl(2, e.level-1)+e.order-2;
  if(value != Nil && T[(i+1)/2-1] == Nil) //给叶子节点复制但是双亲为空
  {
    return ERROR;
  }
  else if(value==Nil &&(T[i*2+1]!=Nil || T[i*2+2]!=Nil)) //给有不为空叶子节点的双亲节点赋空值
  {
    return ERROR;
  }

  T[i] = value;
  return OK;
}

// 返回节点的双亲 如果是根节点则返回空
TElemType Parent(SqBiTree T, TElemType e)
{
  int i;
  if(T[0] == Nil)
  {
    return Nil;
  }

  for(i=1; i<=MAX_TREE_SIZE-1; i++)
  {
    if(T[i] == e)
    {
      return T[(i+1)/2-1];
    }
  }

  return Nil;
}

// 获取节点的左孩子
TElemType LeftChild(SqBiTree T, TElemType e)
{
  int i;
  if(T[0] == Nil)
  {
    return ERROR;
  }

  for(i=0; i<=MAX_TREE_SIZE-1; i++)
  {
    if(T[i] == e)
    {
      return T[i*2+1];
    }
  }

  return Nil;
}


// 获取节点的右孩子
TElemType RightChild(SqBiTree T, TElemType e)
{
  int i;
  if(T[0] == Nil)
  {
    return ERROR;
  }

  for(i=0; i<=MAX_TREE_SIZE-1; i++)
  {
    if(T[i] == e)
    {
      return T[i*2+2];
    }
  }

  return Nil;
}

// 获取节点的左兄弟 若e是T的左孩子或无左兄弟则返回空
TElemType LeftSibling(SqBiTree T, TElemType e)
{
  int i;
  if(T[0] == Nil)
  {
    return ERROR;
  }

  for(i=0; i<=MAX_TREE_SIZE-1; i++)
  {
    if(T[i] == e && i%2 == 0)
    {
      return T[i-1];
    }
  }

  return Nil;
 }


// 获取节点的右兄弟, 若果节点e是右孩子或者无右兄弟则返回空
TElemType RightSibling(SqBiTree T, TElemType e)
{
  int i;
  if(T[0] == Nil)
  {
    return ERROR;
  }

  for(i=0; i<=MAX_TREE_SIZE-1; i++)
  {
    if(T[i] == e && i%2 == 0)
    {
      return T[i+1];
    }
  }

  return Nil;
 }


void PreTraverse(SqBiTree T, int e)
{
  visit(T[e]);
  if(T[2*e+1] != Nil) //左子树不空
  {
    PreTraverse(T, 2*e+1);
  }

  if(T[2*e+2] != Nil) //右子树不空
  {
    PreTraverse(T, 2*e+2);
  }
}

// 先序遍历T
Status PreOrderTraverse(SqBiTree T)
{
  if(!BiTreeEmpty(T))
  {
    PreTraverse(T, 0);
  }
  printf("\n");
  return  OK;
}

void InTraverse(SqBiTree T, int e)
{
  if(T[2*e+1] != Nil) //左子树不为空
  {
    InTraverse(T, 2*e+1);
  }

  visit(e);

  if(T[2*e+2] != Nil) //右子树不为空
  {
    InTraverse(T, 2*e+2);
  }

}


// 中序遍历二叉树
Status InOrderTraverse(SqBiTree T)
{
  if(!BiTreeEmpty(T))
  {
    InTraverse(T, 0);
  }

  printf("\n");
  return OK;
}


void PostTraverse(SqBiTree T, int e)
{
  if(T[2*e+1] != Nil)
  {
    PostTraverse(T, 2*e+1);
  }

  if(T[2*e+2] != Nil)
  {
    PostTraverse(T, 2*e+2);
  }

  visit(T[e]);
}

// 后序遍历二叉树
Status PostOrderTraverse(SqBiTree T)
{
  if(!BiTreeEmpty(T))
  {
    PostTraverse(T, 0);
  }

  printf("\n");
  return OK;
}

//  层序遍历二叉树
void LevelOrderTraverse(SqBiTree T)
{
  int i = MAX_TREE_SIZE -1;
  int j;
  while(T[i] == Nil)
  {
    i--; // 找到最后一个非空节点的序号
  }

  for(j=0; j<=i; j++)
  {
    if(T[j] != Nil)
    {
      visit(T[j]);
    }
  }

  printf("\n");
}

// 逐层 按本层序号输出二叉树
void Print(SqBiTree T)
{
  int j, k;
  Position p;
  TElemType e;

  for(j=1; j<=BiTreeDepth(T); j++)
  {
    printf("第%d层:  ", j);

    for(k=1; k<=pow(2, j-1); k++)
    {
      p.level = j;
      p.order = k;
      e = Value(T, p);
      if (e != Nil)
      {
        printf("%d: %d\n", k, e);
      }
    }

    printf("\n");
  }
}

int main()
{
  Status i;
  Position p;
  TElemType e;
  SqBiTree T;
  InitBiTree(T);
  CreateBiTree(T);
  printf("创建二叉树后树是否为空: %d (1: 是 0:否) 树的深度=%d  \n", BiTreeEmpty(T), BiTreeDepth(T));
  i=Root(T, &e);
  if(i)
  {
    printf("二叉树的根为: %d \n", e);
  }
  else
  {
    printf("数空 无根\n");
  }

  printf("层序遍历二叉: \n");
  LevelOrderTraverse(T);
  printf("前序遍历二叉树: \n");
  PreOrderTraverse(T);
  printf("中序遍历二叉树: \n");
  InOrderTraverse(T);
  printf("后序遍历二叉树: \n");
  PostOrderTraverse(T);
  printf("修改节点的层号3本层序号2. ");
  p.level = 3;
  p.order = 2;
  e = Value(T, p);
  printf("待修改节点的原始值为: %d 新值为: 50", e);
  e = 50;
  Assign(T, p, e);
  printf("前序遍历二叉树: \n");
  PreOrderTraverse(T);
  printf("节点%d的双亲节点为%d", e, Parent(T, e)) ;
  printf("左右孩子分别为:  %d %d ", LeftChild(T, e), RightChild(T, e));
  printf("左右兄弟分别为: %d  %d \n", LeftSibling(T, e), RightSibling(T, e));
  ClearBiTree(T);
  printf("清空二叉树后树是否为空: %d (1: 是 0:否) 树的深度=%d  \n", BiTreeEmpty(T), BiTreeDepth(T));
  i=Root(T, &e);
  if(i)
  {
    printf("二叉树的根为: %d \n", e);
  }
  else
  {
    printf("数空 无根\n");
  }

  return 0;
}
