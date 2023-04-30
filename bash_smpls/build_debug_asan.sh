#!/bin/bash

clang -fsanitize=address -fno-omit-frame-pointer -o app.exe -g main.c -lm