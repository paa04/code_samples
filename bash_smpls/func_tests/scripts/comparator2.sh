#!/bin/bash
file1=$1
file2=$2
v=$3

str1=""
str2=""

if [[ $# -gt 3 ]];then
echo yes
exit 2
fi

if [[ -n "$v" && "$v" != "-v" ]]; then
exit 2
fi

if [[ -f "$file1" && -f "$file2" ]]; then



while read -r string || [[ -n "$string" ]]; do
  str1="$str1""$string"
done < "$file1"


while read -r string || [[ -n "$string" ]]; do
  str2="$str2""$string"
done < "$file2"


str1=${str1#*string:}
str2=${str2#*string:}

#echo $str1 "  " $str2


if [[ "$str1" == "$str2" && -n "$v" ]]; then
  echo "Files are equal"
  exit 0
elif [[ "$str1" == "$str2" ]]; then
  exit 0
elif [[ "$str1" != "$str2" && -n "$v" ]]; then
echo "Files are not equal"
  exit 1
else 
  exit 1
fi

else
exit 2
fi