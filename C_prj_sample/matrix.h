#include "const.h"
#include <stdio.h>
#ifndef MATRIX_H
#define MATRIX_H
int input_size(size_t *, size_t *);
int input_matrix(int matrix[][MAX_COLUMNS], size_t, size_t);
void output_matrix(size_t n, size_t m, int matrix[][MAX_COLUMNS]);
#endif