#!/bin/bash
file1=$1
file1="../data/""$file1"
if [[ $# -gt 1 ]]; then
  exit 3
fi

if [[ -f "../../app.exe" ]]; then
  ../../app.exe <"$file1">"/dev/null"
  code=$?
  if [[ code -eq 1 ]];then
    exit 0
  else
    exit 1
  fi
else
  exit 2
fi