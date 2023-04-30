#include "alg.h"

int main(void)
{
    size_t n, m;
    int status;
    int matrix[MAX_STR][MAX_COLUMNS];
    status = input_size(&n, &m);
    if (status == SUCCESS)
    {
        solve(matrix, n, m);
        output_matrix(n, m, matrix);
    }
    return status;
}