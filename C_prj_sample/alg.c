#include "alg.h"

void solve(int matrix[][MAX_COLUMNS], size_t n, size_t m)
{
    int value = 0;
    for (size_t i = 0; i < m; ++i)

        if (i % 2 == 0)
            for (int j = 0; j < n; ++j)
                matrix[j][i] = ++value;
        else
            for (int j = n - 1; j >= 0; --j)
                matrix[j][i] = ++value;
}
