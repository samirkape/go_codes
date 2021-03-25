#!/bin/bash

unset name

while getopts "c:b:r:d:" opt; do

case $opt in

b)
go build $OPTARG 
mv ${OPTARG%.*} build/
echo "Build Successful"
;;

r)
go build $OPTARG 
result=$?
if [ $result -ne 0 ]
then
    echo " "
    echo "error occured!"
    exit $result
fi
mv `basename ${OPTARG%.*}` build/
echo "Build Successful"
./build/`basename ${OPTARG%.*}` $3
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
