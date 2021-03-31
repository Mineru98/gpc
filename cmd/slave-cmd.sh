#!/bin/bash

if [ $# -eq 1 ]; then
    if [ "list" = "$1" ]; then
        go run slave.go list
    fi
elif [ $# -eq 5 ]; then
    if [ "run" = "$1" ]; then
        go run slave.go run -a $3 -p $5
    elif [ "add" = "$1" ]; then
        echo "Slave Process Add"
    fi
else
    echo "Usage: slave [run] or [list]"
    exit -1
fi