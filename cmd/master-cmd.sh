#!/bin/bash

if [ $# -eq 1 ]; then
    if [ "run" = "$1" ]; then
        go run master.go run
    elif [ "list" = "$1" ]; then
        go run master.go list
    else
        echo "Usage: master [run] or [list]"
    fi
else
    echo "Usage: master [run] or [list]"
    exit -1
fi