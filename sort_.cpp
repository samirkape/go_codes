#include <iostream>
using namespace std;

void sorted(int * arr, int size){
    int patt[3] = {0};
    for(int i=0; i<size; i++){
        if(arr[i] == 0) patt[0] += 1;
        else if(arr[i] == 1) patt[1] += 1;
        else patt[2] += 1;
    }
    int i = 0;
    for(i=0; i<patt[0]; i++) arr[i] = 0;
    for(int j = 0;j<patt[1]; j++) arr[i++] = 1;
    for(int k = 0;k<patt[2]; k++) arr[i++] = 2;
    for(int i=0; i<size; i++) cout<<arr[i]<<" ";
}

int main() {
	//code
   //int arr[] = {7,3,2,5,8}, K = 14;
   int arr[] = {1,0,1,2,1,1,0,0,1,2,1,2,1,2,1,0,0,1,1,2,2,0,0,2,2,2,1,1,1,2,0,0,0,2,0,1,1,1,1,0,0,0,2,2,1,2,2,2,0,2,1,1,2,2,0,2,2,1,1,0,0,2,0,2,2,1,0,1,2,0,0,0,0,2,0,2,2,0,2,1,0,0,2,2}, K = 3;
   int size = sizeof(arr)/sizeof(int);
	sorted(arr,size);
	return 0;
}