#!/bin/bash

this_path="./*.sh"

for file in $this_path; do
  shellcheck "$file"
done

this_path="./func_tests/scripts/*.sh"

for file in $this_path; do
  shellcheck "$file"
done