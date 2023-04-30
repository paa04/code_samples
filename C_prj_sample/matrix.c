#include "matrix.h"

int input_size(size_t *n, size_t *m)
{
    int state = SUCCESS;
    if (scanf("%zu%zu", n, m) != 2 || *n <= 0 || *m <= 0 || *n > MAX_STR || *m > MAX_COLUMNS)
        state = ERROR;
    return state;
}

int input_matrix(int matrix[][MAX_COLUMNS], size_t n, size_t m)
{
    int state = SUCCESS;
    for (size_t i = 0; i < n && state == SUCCESS; ++i)
    {
        for (size_t j = 0; j < m && state == SUCCESS; ++j)
        {
            if (scanf("%d", *(matrix + i) + j) != 1)
            {
                state = ERROR;
            }
        }
    }
    return state;
}

void output_matrix(size_t n, size_t m, int matrix[][MAX_COLUMNS])
{
    for (size_t i = 0; i < n; ++i)
    {
        for (size_t j = 0; j < m; ++j)
        {
            printf("%d ", matrix[i][j]);
        }
        printf("%c",'\n');
    }
}
