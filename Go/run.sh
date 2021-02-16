#!/bin/bash

while getopts "b:r:d:" opt; do

case $opt in

b)
go build $OPTARG 
mv ${OPTARG%.*} build/
echo "Build Successful"
;;

r)
if [ -f build/${OPTARG%.*} ]
then
    echo "Running..."
    ./build/${OPTARG%.*} $3
else
    echo "Build not found..."
    exit -1
fi
;;

d)
if [ -f build/${OPTARG%.*} ]
then
    rm build/${OPTARG%.*}
    echo "Removed!"
fi
;;

*)
#Printing error message
echo "invalid option or argument $OPTARG"
;;

esac
done