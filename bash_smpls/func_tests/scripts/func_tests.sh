#!/bin/bash

not_passed_tests=0
file1="pos_01_in.txt"
file2="pos_01_out.txt"
file3="neg_02_in.txt"
./pos_case.sh "$file1" "$file2"
code=$?
if [[ ! ( $code -eq 0 ) ]];then
  not_passed_tests=$((not_passed_tests+1))
fi

./neg_case.sh "$file3"
code=$?
if [[ ! ( $code -eq 0 ) ]];then
  not_passed_tests=$((not_passed_tests+1))
fi

exit $not_passed_tests