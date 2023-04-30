#include <math.h>
#include <stdio.h>

#define EPS 0.001

double calculation(double x);

int main()
{
    int state;
    double x;
    if (scanf("%lf", &x) == 1)
    {
        state = 0;
    }
    else
        state = 1;
    if (state == 0)
    {
        double ans = calculation(x);
        if (fabs(ans + 1) <= EPS)
        {
            state = 1;
        }
        else
        {
            printf("%.6lf", sqrt(ans));
        }
    }
    return state;
}

double calculation(double x)
{
    if (x < 0)
        return -1;
    int n = 0;
    double ans = 0;
    while (x >= 0)
    {
        ans += x / (++n);
        if (scanf("%lf", &x) == 0)
        {
            ans = -1;
            break;
        }
    }

    return ans;
}