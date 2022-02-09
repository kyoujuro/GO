#!/bin/bash
seq 15 | awk '{print $1%2 ? "a":"b"}' | sort | uniq -c
