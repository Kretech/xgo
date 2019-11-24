#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <time.h>
#include <sys/time.h>

const int n = 702410;
int a[n];

double dosth()
{
	struct timeval begin;
	struct timeval end;
	gettimeofday(&begin, NULL);
	for (int i = 0; i < n; i++)
	{
		a[i] = a[i] * i + i;
	}
	gettimeofday(&end, NULL);
	return end.tv_sec - begin.tv_sec + (end.tv_usec - begin.tv_usec) *1.0/ 1000000;
}

int main()
{
	int id = getpid();
	for (int i = 0; i < n; i++)
	{
		a[i] = i;
	}

	char a[32];
	int c = 0;
	while (~scanf("%s", a))
	{
		double t = dosth();
		printf("%d:|%s|%lu|t:%f\n", id, a, strlen(a), t);
		fflush(stdout);
	}
	return 0;
}
