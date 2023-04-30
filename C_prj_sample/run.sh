#!/bin/bash

gcc -c *c
gcc -g -o app_debug.exe *o
gcc -o app.exe *o