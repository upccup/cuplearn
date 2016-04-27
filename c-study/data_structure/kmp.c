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
typedef int ElemType;

typedef char String[MAXSIZE+1];


// 生成一个其值等于char的串T
Status StrAssign(String T, char *chars)
{
  int i;
  if(strlen(chars) > MAXSIZE)
    return ERROR;

  T[0] = strlen(chars);
  for(i=1; i<= T[0]; i++)
  {
    T[i] = *(chars+i-1);
  }
  return OK;
}

// 清空字符串T
Status ClearString(String S)
{
  S[0]=0;
  return OK;
}


// 输出字符串T
void StrPrint(String T)
{
  int i;
  for(i=1; i<= T[0]; i++)
  {
    printf("%c  ", T[i]);
  }
  printf("\n");
}

// 输出next数组
void NextPrint(int next[], int length)
{
  int i;
  for(i=1; i<= length; i++)
  {
    printf("%d  ", next[i]);
  }
  printf("\n");
}

// 获取串中元素的个数
int StrLength(String S)
{
  return S[0];
}

// 朴素的模式匹配
int Index(String S, String T, int pos)
{
  int i = pos;
  int j = 1;
  while(i <= S[0] && j <= T[0])
  {
    if (S[i] == T[j])
    {
      ++i;
      ++j;
    }
    else
    {
      i = i-j+2;
      j = 1;
    }
  }

  if (j > T[0])
    return i-T[0];

  return 0;
}

// 通过计算返回子串的Next数组
void getNext(String T, int *next)
{
  int i, j;
  i = 1;
  j = 0;
  next[1] = 0;

  while(i<T[0])
  {
    if(j==0 || T[i] == T[j])
    {
      ++i;
      ++j;
      next[i] = j;
    }
    else
    {
      j = next[j];
    }
  }
}

// KMP模式匹配查找
int IndexKMP(String S, String T, int pos)
{
  int i = pos;
  int j = 1;
  int next[255];
  getNext(T, next);

  while(i <= S[0] && j <= T[0])
  {
    if (j == 0 || S[i] == T[j])
    {
      ++i;
      ++j;
    }
    else
    {
      j = next[j];
    }
  }

  if (j > T[0])
    return i - T[0];

  return 0;
}


// 获取模式串T的next函数修正值并存入数据组nextval
void getNextval(String T, int *nextval)
{
  int i, j;
  i = 1;
  j = 0;
  nextval[1] = 0;

  while (i < T[0])
  {
    if(j==0 || T[i] == T[j])
    {
      ++i;
      ++j;
      if (T[i] != T[j])
      {
        nextval[i] = j;
      }
      else
      {
        nextval[i] = nextval[j];
      }
    }
    else
    {
      j = nextval[j];
    }
  }
}

int IndexKMP1(String S, String T, int pos)
{
  int i = pos;
  int j = 1;
  int next[255];
  getNextval(T, next);
  while(i <= S[0] && j <= T[0])
  {
    if(j == 0  || S[i] == T[j])
    {
      ++i;
      ++j;
    }
    else
    {
      j = next[j];
    }
  }

  if(j > T[0])
    return i - T[0];

  return 0;
}

int main()
{
  int i, *p;
  String s1, s2;

  StrAssign(s1, "abcdex");
  printf("子串为:     ");
  StrPrint(s1);
  i = StrLength(s1);
  p = (int*)malloc((i+1)*sizeof(int));
  getNext(s1, p);
  printf("Next数组为:  ");
  NextPrint(p, StrLength(s1));
  printf("\n");

  StrAssign(s1, "abcabx");
  printf("子串为:     ");
  StrPrint(s1);
  i = StrLength(s1);
  p = (int*)malloc((i+1)*sizeof(int));
  getNext(s1, p);
  printf("Next数组为:  ");
  NextPrint(p, StrLength(s1));
  printf("\n");

  StrAssign(s1, "ababaaaba");
  printf("子串为:     ");
  StrPrint(s1);
  i = StrLength(s1);
  p = (int*)malloc((i+1)*sizeof(int));
  getNext(s1, p);
  printf("Next数组为:  ");
  NextPrint(p, StrLength(s1));
  printf("\n");

  StrAssign(s1, "aaaaaaaaaab");
  printf("子串为:     ");
  StrPrint(s1);
  i = StrLength(s1);
  p = (int*)malloc((i+1)*sizeof(int));
  getNext(s1, p);
  printf("Next数组为:  ");
  NextPrint(p, StrLength(s1));
  printf("\n");

  StrAssign(s1, "ababaaaba");
  printf("子串为:     ");
  StrPrint(s1);
  i = StrLength(s1);
  p = (int*)malloc((i+1)*sizeof(int));
  getNext(s1, p);
  printf("Next数组为:  ");
  NextPrint(p, StrLength(s1));
  printf("\n");
  getNextval(s1, p);
  printf("NextVal为:   ");
  NextPrint(p, StrLength(s1));
  printf("\n");

  StrAssign(s1, "aaaaaaaaaab");
  printf("子串为:     ");
  StrPrint(s1);
  i = StrLength(s1);
  p = (int*)malloc((i+1)*sizeof(int));
  getNext(s1, p);
  printf("Next数组为:  ");
  NextPrint(p, StrLength(s1));
  printf("\n");
  getNextval(s1, p);
  printf("NextVal为:   ");
  NextPrint(p, StrLength(s1));
  printf("\n");

  StrAssign(s1, "000000000000000000000000000000000000000000000000000000000000000001");
  printf("主串为:   ");
  StrPrint(s1);
  StrAssign(s2, "00000000001");
  printf("子串为:     ");
  StrPrint(s2);
  printf("\n");
  printf("主串和子串在%d个字符处首次匹配(朴素匹配算法) \n", Index(s1, s2, 1));
  printf("主串和子串在%d个字符处首次匹配(KMP算法) \n", IndexKMP(s1, s2, 1));
  printf("主串和子串在%d个字符处首次匹配(KMP改良算法) \n", IndexKMP1(s1, s2, 1));

  return 0;
}
