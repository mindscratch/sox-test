#!/bin/bash

wavfile="$1"
outfile="$2"

sox "$wavfile" -r 8000 -t ul "$outfile"
