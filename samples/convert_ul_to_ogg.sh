#!/bin/bash

input="$1"
output="$2"

sox -r 8000 -t ul "$input" -C 2 "$output"
