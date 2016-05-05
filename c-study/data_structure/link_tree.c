#include <stdio.h>
#include <string.h>
#include <stdlib.h>


#include <math.h>
#include <time.h>


#define OK 1
#define ERROR 0
#define TRUE 1
#define FALSE 0

#define MAXSIZE 100

typedef int Status;

int index1=1;
typedef char String[24];
String str;

Status StrAssign(String T, char *chars)
{
  int i;
  if (strlen(chars) > MAXSIZE)
  {
    return ERROR;
  }
  else
  {
    T[0] = strlen(chars);
    printf("%d  \n", T[0]);
    for(i = 1; i <= T[0]; i++)
    {
      printf("%c \n", *(chars + i - 1));
      T[i] = *(chars + i - 1);
      // printf("%c ", T[i]);
    }
    return OK;
  }
}

typedef char TElemType;
TElemType Nil = ' '; // 字符型以空格符为空

Status visit(TElemType e)
{
  printf("%c  ", e);
  return OK;
}

typedef struct BiTNode //节点结构
 {
   TElemType data; // 节点数据
   struct BiTNode *lchild, *rchild;
 } BiTNode, *BiTree;


// 构造空二叉树
Status InitBiTree(BiTree *T)
{
  *T = NULL;
  return OK;
}

//销毁二叉树
void DestroyBiTree(BiTree *T)
{
  if(*T)
  {
    if((*T)->lchild) // 存在左孩子
    {
      DestroyBiTree(&(*T)->lchild);
    }
    else //存在右孩子
    {
      DestroyBiTree(&(*T)->rchild);
    }

    free(*T);
    *T = NULL;
  }
}

// 构造二叉树
void CreateBiTree(BiTree *T)
{
  TElemType ch;
  ch = str[index1++];

  if(ch == '#')
  {
     *T = NULL;
  }
  else
  {
    *T = (BiTree)malloc(sizeof(BiTNode));
    if(!*T)
    {
      exit(OVERFLOW);
    }

    (*T)->data = ch; //生成根节点
    CreateBiTree(&(*T)->lchild); //构造左子树
    CreateBiTree(&(*T)->rchild); //构造右子树
  }
}

// 判断二叉树是否为空
Status BiTreeEmpty(BiTree T)
{
  if(T)
  {
    return FALSE;
  }
  else
  {
    return TRUE;
  }
}

#define ClearBiTree DestroyBiTree

// 获取二叉树的深度
int BiTreeDepth(BiTree T)
{
  int i, j;
  if(!T)
  {
    return 0;
  }

  if(T->lchild)
  {
    i = BiTreeDepth(T->lchild);
  }
  else
  {
    i = 0;
  }

  if(T->rchild)
  {
    j = BiTreeDepth(T->rchild);
  }
  else
  {
    j = 0;
  }

  return i>j?i+1:j+1;
}

// 获取二叉树的根节点
TElemType Root(BiTree T)
{
  if(BiTreeEmpty(T))
  {
    return Nil;
  }
  else
  {
   return T->data;
  }
}

// 返回指点位置节点的值
TElemType Value(BiTree p)
{
  return p->data;
}

// 给指定节点赋值
void Assign(BiTree p, TElemType value)
{
  p->data = value;
}

// 前序递归遍历二叉树
void PreOrderTraverse(BiTree T)
{
  if(T == NULL)
  {
    return;
  }

  printf("%c  ", T->data);
  PreOrderTraverse(T->lchild); // 先递归遍历左子树
  PreOrderTraverse(T->rchild); // 在递归遍历右子树
}


// 中序遍历二叉树
void InOrderTraverse(BiTree T)
{
  if(T == NULL)
  {
    return;
  }

  InOrderTraverse(T->lchild); //中序遍历左子树
  printf("%c  ", T->data); // 显示节点数据
  InOrderTraverse(T->rchild); //中序遍历右子树
}

// 后序遍历二叉树
void PostOrderTraverse(BiTree T)
{
  if(T == NULL)
  {
    return;
  }

  PostOrderTraverse(T->lchild); // 先后序遍历左子树
  PostOrderTraverse(T->rchild); // 在后序遍历右子树
  printf("%c   ", T->data); // 显示节点数据
}

int main()
{
  int i;
  BiTree T;
  TElemType e1;
  InitBiTree(&T);

  StrAssign(str,"ABDH#K###E##CFI###G#J##");
  // StrAssign(str, "AFSFSFFS$###FISDD##GJ#O##");

  CreateBiTree(&T);

  printf("构造二叉树后树是否为空? %d (1: 是, 0:否) 树的深度= %d \n", BiTreeEmpty(T), BiTreeDepth(T));
  e1 = Root(T);
  printf("二叉树的根为: %c\n", e1);

  printf("\n 前序遍历二叉树: ");
  PreOrderTraverse(T);

  printf("\n 中序遍历二叉树: ");
  InOrderTraverse(T);

  printf("\n 后序遍历二叉树: ");
  PostOrderTraverse(T);

  printf("\n");
  ClearBiTree(&T);
  printf("清空二叉树后树是否为空? %d (1: 是, 0:否) 树的深度= %d \n", BiTreeEmpty(T), BiTreeDepth(T));

  i = Root(T);
  if (!i)
  {
    printf("树空 无根 \n");
  }
  return 0;
}
