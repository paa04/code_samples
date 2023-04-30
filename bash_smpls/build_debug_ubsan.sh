#!/bin/bash

clang -fsanitize=undefined -fno-omit-frame-pointer -o app.exe -g main.c -lm