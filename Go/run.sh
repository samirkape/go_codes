#!/bin/bash

go build $1
result=$?
if [ $result -ne 0 ]
then
    echo " "
    echo "error occured!"
    exit $result
fi
mv `basename ${1%.*}` build/
echo "Build Successful"
./build/`basename ${1%.*}` $2

