#include<stdio.h>
#include<math.h>
#include<stdlib.h>
#include<time.h>
int choose1,choose,row,line,times,be,i,reI,ReI;
int array2[]={};
char value;
char array[3][3]={ ' ',' ',' ' ,' ' ,' ' ,' ' ,' ' ,' ' ,' ' };
int Choose_game()
{
	//开始界面// 
	printf("************* \n 开 始 游 戏 \n************* \n");
	printf("请选择游戏：\n\n");
	printf("1: 三子棋\n2: out\n");
	scanf("%d",&choose1);
	if(choose1==2)
	{
		return 3;
	}
	printf("请选择模式：\n1: PVP  2: PVE\n");
	scanf("%d",&choose);
	return choose; 
}

//检查落点重复// 
int Inspection(int j,int row,int line)
{
	times=10*row+line;
	array2[j]=times;
	for(int m=0;m<=j-1;m++)
	{
		if(array2[m]==times)
		{
			return 1;
			break;
		}
	}
	return 0;
}

int Start(int i,int row,int line)
{
	if(i%2==0)
		value='#';
	else
		value='*';
	printf("row=%d,line=%d\n",row,line);
	if(row==1)
	{
		switch (line)
		{
			case 1:array[0][0]=value;break;
			case 2:array[0][1]=value;break;
			case 3:array[0][2]=value;break;
		}
	}
	if(row==2)
	{
		switch (line)
		{
			case 1:array[1][0]=value;break;
			case 2:array[1][1]=value;break;
			case 3:array[1][2]=value;break;
		}
	}
	if(row==3)
	{
		switch (line)
		{
			case 1:array[2][0]=value;break;
			case 2:array[2][1]=value;break;
			case 3:array[2][2]=value;break;
		}
	}
	printf("| %c | %c | %c | \n ------------ \n| %c | %c | %c | \n ------------ \n| %c | %c | %c | \n\n",
	array[0][0],array[0][1],array[0][2],array[1][0],array[1][1],
	array[1][2],array[2][0],array[2][1],array[2][2]);

}

int End()
{
	for(int n=0;n<3;n++)
	{
		//横竖判定// 
		if(((array[n][0]==array[n][1]&&array[n][2]==array[n][1])&&array[n][1]=='#')||((array[0][n]==array[1][n]&&array[2][n]==array[1][n])&&array[1][n]=='#'))
		{
			return 1;
			break;
		}
		if(((array[n][0]==array[n][1]&&array[n][2]==array[n][1])&&array[n][1]=='*')||((array[0][n]==array[1][n]&&array[2][n]==array[1][n])&&array[1][n]=='*'))
		{
			return 2;
			break;
		}
	}
	//对角线判定// 
	if((array[0][0]==array[1][1]&&array[1][1]==array[2][2])||(array[2][0]==array[1][1]&&array[1][1]==array[0][2]))
	{
		if((array[0][0]==array[1][1]&&array[1][1]=='#')||(array[2][0]==array[1][1]&&array[2][0]=='#'))
		{
			//玩家1赢// 
			return 1;
		}
		if((array[0][0]==array[1][1]&&array[1][1]=='*')||(array[2][0]==array[1][1]&&array[2][0]=='*'))
		{
			//玩家2 or 电脑赢// 
			return 2; 
		} 
	}
}


void Show()
{
	//棋盘展示// 
	printf("|   |   |   | \n ------------ \n|   |   |   | \n ------------ \n|   |   |   | \n");
}
	

int main()
{ 
	int q,p;
  	tma:
	be=Choose_game();
	switch (be)
	{
		case 1:
			Show();
			printf("玩家一: #\n玩家二: *\n");
			for(i=0;i<9;i++)
			{
				tmp:
				printf("--->");
				printf("请玩家%d选择落子的行数(1，2，3)：\n",i%2+1);
				scanf("%d",&row);
				if(row<0||row>3)
				{
					printf("所输入的行数超出所给范围！\n");
					goto tmp;
				}
				tmq:
				printf("--->");
				printf("请玩家%d选择落子的列数(1，2，3): \n",i%2+1);
				scanf("%d",&line);
				if(line<0||line>3)
				{
					printf("所输入的行数超出所给范围！\n");
					goto tmq;
				}
				reI=Inspection(i,row,line);
				if(reI==1)
				{
					printf("落子重复！请重新落子\n");
					goto tmp;
				}
				Start(i,row,line);
				int N=End();
				if(N==1)
				{
					printf("玩家1 获胜！...\n");
					break;
				}
				if(N==2)
				{
					printf("玩家2 获胜！...\n");
					break;
				}
				if(i==7)
				{
					printf("平局！\n");
					break;
				}
			}	
			break;
		case 2:
			Show();
			printf("玩家一: #\n电脑: *\n");
			for(i=0;i<9;i++)
			{
				if(i%2==0)
				{
					tmP:
					printf("请玩家选择落子的行数(1，2，3)：\n");
					scanf("%d",&row);
					if(row<0||row>3)
					{
						printf("所输入的行数超出所给范围！\n");
						goto tmP;
					}
					tmQ:
					printf("请玩家选择落子的列数(1，2，3): \n");
					scanf("%d",&line);
					if(line<0||line>3)
					{
						printf("所输入的行数超出所给范围！\n");
						goto tmQ;
					}
					reI=Inspection(i,row,line);
					if(reI==1)
					{
						printf("落子重复！请重新落子\n");
						goto tmP;
					}
				}
				else
				{
					printf("电脑思考中...\n");
					tmc:
					//生成随机数// 
					srand(time(0));
					row=1+rand()%3;
					line=1+rand()%3;
					//检查落点重复// 
					ReI=Inspection(i,row,line);
					if(ReI==1)
					{
						goto tmc;
					}
					printf("电脑已落子...\n\n");
				}
				Start(i,row,line);
				int N=End();
				if(N==1)
				{
					printf("玩家1 获胜！...\n");
					break;
				}
				if(N==2)
				{
					printf("电脑获胜，你可以回家修理地球了！...\n");
					break;
				}
				if(i==7)
				{
					printf("平局！");
					break;
				}
			}
			break;
		case 3:
			{
				printf("已退出...");
				break;
			}
		default :
			printf("未知输入！");
			goto tma;
			break; 
    }
    return 0;
}

