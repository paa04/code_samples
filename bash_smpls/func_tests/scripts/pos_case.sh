#!/bin/bash
file1=$1
file2=$2
file1="../data/""$file1"
file2="../data/""$file2"
touch ../data/fact_ans.txt
fact_ans="../data/fact_ans.txt"
#echo $file1
if [[ $# -gt 2 ]]; then
  exit 3
fi
#echo $fact_ans
if [[ -f "../../app.exe" ]]; then
  ../../app.exe <"$file1" > "$fact_ans"
  ./comparator2.sh "$fact_ans" "$file2"
  code=$?
  if [[ $code -eq 0 ]];then
    exit 0
  else
    exit 1
  fi
else
  exit 2
fi

