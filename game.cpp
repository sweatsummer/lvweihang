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
	//��ʼ����// 
	printf("************* \n �� ʼ �� Ϸ \n************* \n");
	printf("��ѡ����Ϸ��\n\n");
	printf("1: ������\n2: out\n");
	scanf("%d",&choose1);
	if(choose1==2)
	{
		return 3;
	}
	printf("��ѡ��ģʽ��\n1: PVP  2: PVE\n");
	scanf("%d",&choose);
	return choose; 
}

//�������ظ�// 
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
		//�����ж�// 
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
	//�Խ����ж�// 
	if((array[0][0]==array[1][1]&&array[1][1]==array[2][2])||(array[2][0]==array[1][1]&&array[1][1]==array[0][2]))
	{
		if((array[0][0]==array[1][1]&&array[1][1]=='#')||(array[2][0]==array[1][1]&&array[2][0]=='#'))
		{
			//���1Ӯ// 
			return 1;
		}
		if((array[0][0]==array[1][1]&&array[1][1]=='*')||(array[2][0]==array[1][1]&&array[2][0]=='*'))
		{
			//���2 or ����Ӯ// 
			return 2; 
		} 
	}
}


void Show()
{
	//����չʾ// 
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
			printf("���һ: #\n��Ҷ�: *\n");
			for(i=0;i<9;i++)
			{
				tmp:
				printf("--->");
				printf("�����%dѡ�����ӵ�����(1��2��3)��\n",i%2+1);
				scanf("%d",&row);
				if(row<0||row>3)
				{
					printf("���������������������Χ��\n");
					goto tmp;
				}
				tmq:
				printf("--->");
				printf("�����%dѡ�����ӵ�����(1��2��3): \n",i%2+1);
				scanf("%d",&line);
				if(line<0||line>3)
				{
					printf("���������������������Χ��\n");
					goto tmq;
				}
				reI=Inspection(i,row,line);
				if(reI==1)
				{
					printf("�����ظ�������������\n");
					goto tmp;
				}
				Start(i,row,line);
				int N=End();
				if(N==1)
				{
					printf("���1 ��ʤ��...\n");
					break;
				}
				if(N==2)
				{
					printf("���2 ��ʤ��...\n");
					break;
				}
				if(i==7)
				{
					printf("ƽ�֣�\n");
					break;
				}
			}	
			break;
		case 2:
			Show();
			printf("���һ: #\n����: *\n");
			for(i=0;i<9;i++)
			{
				if(i%2==0)
				{
					tmP:
					printf("�����ѡ�����ӵ�����(1��2��3)��\n");
					scanf("%d",&row);
					if(row<0||row>3)
					{
						printf("���������������������Χ��\n");
						goto tmP;
					}
					tmQ:
					printf("�����ѡ�����ӵ�����(1��2��3): \n");
					scanf("%d",&line);
					if(line<0||line>3)
					{
						printf("���������������������Χ��\n");
						goto tmQ;
					}
					reI=Inspection(i,row,line);
					if(reI==1)
					{
						printf("�����ظ�������������\n");
						goto tmP;
					}
				}
				else
				{
					printf("����˼����...\n");
					tmc:
					//���������// 
					srand(time(0));
					row=1+rand()%3;
					line=1+rand()%3;
					//�������ظ�// 
					ReI=Inspection(i,row,line);
					if(ReI==1)
					{
						goto tmc;
					}
					printf("����������...\n\n");
				}
				Start(i,row,line);
				int N=End();
				if(N==1)
				{
					printf("���1 ��ʤ��...\n");
					break;
				}
				if(N==2)
				{
					printf("���Ի�ʤ������Իؼ���������ˣ�...\n");
					break;
				}
				if(i==7)
				{
					printf("ƽ�֣�");
					break;
				}
			}
			break;
		case 3:
			{
				printf("���˳�...");
				break;
			}
		default :
			printf("δ֪���룡");
			goto tma;
			break; 
    }
    return 0;
}

