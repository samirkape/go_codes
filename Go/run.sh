#!/bin/bash

if [ ! -f build/${1%.*} ]
then
    echo "Building..."
    go build $1 && mv ${1%.*} ./build/ && ./build/${1%.*} $2
else
    echo "Running..."
    ./build/${1%.*} $2
fi