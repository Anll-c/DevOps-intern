#!/usr/bin/env bash

log_count=$(ls *.log | wc -w)
echo ".log: $log_count"
total_size=0

for (( i=1; i<=log_count; i++ )); do
  size=$(du *.log | awk -v row="$i" 'NR==row {print $1}')
  total_size=$((total_size + size))
done

echo "total size::: $total_size"