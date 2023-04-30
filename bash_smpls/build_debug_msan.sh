#!/bin/bash

clang -fsanitize=memory -fPIE -pie -fno-omit-frame-pointer -o app.exe -g main.c -lm